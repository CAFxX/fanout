package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int32; F2 int }
type GenS5 struct { F0 int32; F1 int32; F2 bool }
func genFn0(p0 string, p1 uint16, p2 string) int32 {
    if ((p1 % ((p1) | 1)) == p1) {
        p2 = (p2 + p0)
        var v7 uint64 = 142
        _ = v7
    }
    if false {
        if (!(false || true)) {
            for fi8 := 0; fi8 < 8; fi8++ {
                switch 1 {
                case 9223372036854775807:
                    var v9 uint8 = 61
                    _ = v9
                case 0:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v11 := GenS1{F0: fi10, F1: 4294967295}
                        fmt.Println(p2, p0)
                        var v12 uint16 = p1
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case 127:
                    var v13 int64 = 9223372036854775807
                    _ = v13
                }
                p0 = "hello"
                var v14 uint = 94
                _ = fi8
                _ = v14
            }
        } else {
            var v15 int64 = 1
            var v16 bool = (p1 > 287)
            if v16 {
                switch 42 {
                case -9223372036854775808:
                    v17 := []int{-1, 128, 0}
                    _ = v17
                }
                if ((((v15 >> ((2) & 63)) >= -1) && (v16 && (v15 <= -425))) || false) {
                    v18 := GenS2{F0: -1, F1: 1, F2: 255}
                    _ = v18
                } else {
                    v19 := GenS3{F0: 117, F1: p1}
                    var v20 int64 = (v15 ^ v15)
                    _ = v19
                    _ = v20
                }
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 int16 = -32768
                    var v23 uint64 = 939
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            _ = v15
            _ = v16
        }
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            p2 = p0
            _ = fi25
        }
        var v26 int64 = -1
        var v27 int64 = 9223372036854775807
        _ = fi24
        _ = v26
        _ = v27
    }
    switch 127 {
    case -2147483648:
        p0 = (p0 + ((((("" + ((p0 + "") + p0)) + p0) + p2) + p2) + p2))
    case 127:
        v28 := GenS2{F0: 0, F1: 255, F2: 99}
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    return 849
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]bool)
    var v3 uint32 = 65535
    var v4 string = (("foo bar" + "") + "a")
    var v5 uint32 = 1
    var v6 uint32 = (^(v5))
    if false {
        v29 := GenS4{F0: 178, F1: 2147483647, F2: (-(v1))}
        _ = v29
    } else {
        v5 += (^(v5))
        for fi30 := 0; fi30 < 3; fi30++ {
            v5 += 42
            var v31 float64 = ((float64((-2.25 / 1e10)) / 0.1) / float64((0.1 / 3.14159)))
            if false {
                var v32 float32 = float32(v31)
                switch v1 {
                case 317:
                    v33 := []int{(fi30 << ((4294967295) & 63)), (fi30 | (v1 ^ (fi30 * fi30)))}
                    _ = v33
                case 9223372036854775807:
                    v34 := []int{fi30, 0, -2147483648}
                    _ = v34
                case 32767:
                    v35 := []int{(fi30 | (v1 | (-(v1)))), (^(fi30))}
                    _ = v35
                }
                _ = v32
            }
            _ = fi30
            _ = v31
        }
        v36 := GenS3{F0: v6, F1: 308}
        _ = v36
    }
    v3 = (v3 >> ((65535) & 63))
    fmt.Println(v4, v6, v5)
    v5 = (v6 << ((255) & 63))
    v2[v6] = true
    if ((1.5 != -2.25) && false) {
        fmt.Println(v1, v3, v1)
        var v37 uint32 = (v3 ^ (v6 << ((42) & 63)))
        _ = v37
    } else {
        var v38 uint32 = v3
        switch v4 {
        case "":
            var v39 string = "hello"
            _ = v39
        case "hello":
            switch (v1 | 0) {
            case 562:
                v40 := GenS1{F0: v1, F1: 68}
                _ = v40
            case 32768:
                for fi41 := 0; fi41 < 12; fi41++ {
                    v2[v38] = (393 > 42)
                    v42 := GenS5{F0: 1, F1: 42, F2: true}
                    _ = fi41
                    _ = v42
                }
            default:
                if (uint16(v3) < 1) {
                    v6 = v38
                } else {
                    delete(v2, (v38 * (v5 << ((4294967295) & 63))))
                }
            }
        case "a":
            v38 -= (^(v6))
        }
        _ = v38
    }
    v43 := GenS5{F0: -2147483648, F1: -2147483648, F2: false}
    var v44 uint64 = 42
    var v45 float32 = 1e10
    v45 += (-2.25 + 1e10)
    v4 = v4
    for fi46 := 0; fi46 < 5; fi46++ {
        fmt.Println(v44, fi46)
        switch -9223372036854775808 {
        case 32768:
            v3 -= (v5 / ((42) | 1))
        case 128:
            if (true || false) {
                v1 += (v1 + (-(fi46)))
                for fi47 := 0; fi47 < 10; fi47++ {
                    var v48 float32 = v45
                    fmt.Println(v48, fi47)
                    _ = fi47
                    _ = v48
                }
                v45 = 1.5
            }
        case -439:
            var v49 int8 = 100
            _ = v49
        default:
            var v50 string = ""
            _ = v50
        }
        _ = fi46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
