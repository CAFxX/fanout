package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int }
func genFn0(p0 int64, p1 int) uint64 {
    fmt.Println(p1, p1)
    p1 = (p1 & p1)
    var v5 int16 = -742
    var v6 float32 = float32(0.0)
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() int16 {
    var v8 uint8 = 255
    v9 := GenS1{F0: 2, F1: v8, F2: 931}
    v8 += (v8 * 251)
    _ = v8
    _ = v9
    return 794
}

func main() {
    v1 := []uint32{2147483648, 179, 2, 65535}
    var v2 int = -824
    v3 := []string{("hello" + (("" + "foo bar") + "a")), "", ""}
    v4 := make(map[uint32]int)
    if false {
        v10 := GenS2{F0: 0, F1: -1, F2: -32768}
        _ = v10
    }
    fmt.Println(v2)
    v4[0] = (v2 % ((v2) | 1))
    v2 = (v2 >> ((0) & 63))
    if true {
        var v11 uint64 = 4294967295
        v12 := GenS1{F0: v11, F1: 255, F2: 32767}
        v11 = genFn0(228, v2)
        _ = v11
        _ = v12
    }
    switch "hello" {
    case "hello":
        if (!true) {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS3{F0: (52 > 165)}
                _ = fi13
                _ = v14
            }
        } else {
            if ((2147483647 < 1) || (true || false)) {
                v15 := GenS4{F0: 18446744073709551615}
                _ = v15
            }
            if len(v1) > 0 {
                v1[0] = 424
            }
        }
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v16 uint = 65535
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 string = "\n\t"
        fmt.Println(fi17, v18)
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        if (true || (42 >= 128)) {
            v16 = (v16 | 18446744073709551615)
            var v20 uint16 = 697
            if (false || true) {
                fmt.Println(v2, fi19)
                fmt.Println(fi19, v16)
            }
            _ = v20
        } else {
            v21 := GenS5{F0: 128}
            if false {
                v16 = (v16 % (((v16 - v16)) | 1))
            }
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int = v2
                v2 = 127
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        v24 := GenS4{F0: 1}
        v4[2147483648] = (v2 << ((v16) & 63))
        _ = fi19
        _ = v24
    }
    for fi25 := 0; fi25 < 12; fi25++ {
        v16 = (v16 | v16)
        if (42 == 4294967295) {
            var v26 uint16 = 42
            _ = v26
        } else {
            if len(v1) > 0 {
                v1[0] = 2147483648
            }
        }
        v27 := GenS6{F0: 100}
        _ = fi25
        _ = v27
    }
    var v28 uint16 = 65535
    var v29 int8 = 100
    _ = len(v3)
    if true {
        v4[285] = 32768
        switch 127 {
        case 127:
            for fi30 := 0; fi30 < 2; fi30++ {
                var v31 int64 = 0
                _ = fi30
                _ = v31
            }
        case 32768:
            var v32 uint64 = uint64(v16)
            _ = v32
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            if true {
                v3 = append(v3, "a")
                var v34 int64 = 236
                var v35 bool = true
                _ = v34
                _ = v35
            } else {
                for fi36 := 0; fi36 < 14; fi36++ {
                    v2 += 127
                    _ = fi36
                }
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 uint16 = 2
                    var v39 int32 = -2147483648
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            var v40 int64 = -398
            if (190 != 127) {
                fmt.Println(v2, v29, v28)
                fmt.Println(fi33)
                v41 := GenS6{F0: 0}
                _ = v41
            }
            _ = fi33
            _ = v40
        }
    }
    switch "" {
    case "hello":
        if (!(32767 > 255)) {
            fmt.Println(v16)
        }
    case "":
        v16 = v16
    case "foo bar":
        v42 := GenS6{F0: v29}
        _ = v42
    }
    for fi43 := 0; fi43 < 5; fi43++ {
        v2 -= fi43
        v44 := GenS7{F0: 1, F1: 2, F2: 270}
        _ = len(v3)
        _ = fi43
        _ = v44
    }
    v28 = v28
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
