### Build
```sh
make kernel initramfs
```

### Run QEMU
```sh
qemu-system-x86_64 -kernel build/kernel -initrd build/initramfs.gz -nographic -append "console=ttyS0"
```
