package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint8, p2 int16) string {
    v4 := []int{32767, 4294967295}
    var v5 float64 = 0.1
    v4 = append(v4, 97)
    var v6 uint64 = 42
    v7 := GenS1{F0: 255, F1: 173, F2: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return "a"
}

func main() {
    var v1 int8 = 42
    var v2 uint32 = 2147483648
    v3 := make(map[uint32]int)
    v8 := []int{1, 150, -2147483648}
    var v9 int32 = 42
    var v10 uint8 = 42
    var v11 uint64 = 1
    v12 := GenS1{F0: 42, F1: v9, F2: v2}
    if len(v8) > 0 {
        v8[0] = 1
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        switch fi13 {
        case 659:
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint8 = (^(v10))
                var v16 int64 = 9223372036854775807
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            if true {
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 string = ""
                    _ = fi17
                    _ = v18
                }
            } else {
                v9 += v9
                for fi19 := 0; fi19 < 8; fi19++ {
                    if (1e10 > (1.5 * 3.14159)) {
                        var v20 uint8 = 98
                        _ = v20
                    }
                    if true {
                        var v21 uint32 = v2
                        _ = v21
                    }
                    var v22 int = fi13
                    switch 390 {
                    case 2147483647:
                        v2 -= v2
                    case -245:
                        var v23 int32 = (v9 % (((v9 * v9)) | 1))
                        _ = v23
                    }
                    var v24 float32 = 0.1
                    _ = fi19
                    _ = v22
                    _ = v24
                }
                v11 = (v11 >> ((0) & 63))
            }
        }
        v25 := GenS1{F0: 911, F1: 740, F2: (^(v2))}
        for fi26 := 0; fi26 < 5; fi26++ {
            v10 = 0
            fmt.Println(v2, v1)
            _ = fi26
        }
        _ = fi13
        _ = v25
    }
    switch "a" {
    case "a":
        v9 = v9
    case "hello":
        if (!true) {
            var v27 int32 = v9
            var v28 int16 = 42
            var v29 float64 = (1e10 - -2.25)
            _ = v27
            _ = v28
            _ = v29
        } else {
            fmt.Println(v11, v11)
            for fi30 := 0; fi30 < 11; fi30++ {
                if false {
                    var v31 uint64 = v11
                    var v32 int32 = (v9 % ((v9) | 1))
                    var v33 uint16 = 2
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    if ("" == ("hello" + "\n\t")) {
                        v1 -= v1
                    }
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v10 = 255
                        var v35 int32 = 42
                        if len(v8) > 0 {
                            v8[0] = (fi30 ^ (fi34 * 967))
                        }
                        _ = fi34
                        _ = v35
                    }
                }
                if ((-537 == 981) && (65535 > 253)) {
                    var v36 uint64 = v11
                    fmt.Println(fi30)
                    _ = v36
                } else {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    fmt.Println(v11, v2, fi30)
                }
                _ = fi30
            }
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 uint8 = v10
                _ = fi37
                _ = v38
            }
        }
    default:
        for fi39 := 0; fi39 < 10; fi39++ {
            v40 := GenS2{F0: 255}
            _ = fi39
            _ = v40
        }
    }
    if (1 != 65535) {
        v9 = 42
        v41 := GenS3{F0: (v11 << ((18446744073709551615) & 63))}
        _ = v41
    } else {
        var v42 int8 = (v1 / (((v1 & int8(v9))) | 1))
        var v43 int = 781
        _ = v42
        _ = v43
    }
    var v44 uint64 = (v11 + v11)
    if false {
        for fi45 := 0; fi45 < 1; fi45++ {
            v46 := GenS1{F0: 896, F1: v9, F2: 2147483647}
            if len(v8) > 0 {
                v8[0] = fi45
            }
            _ = fi45
            _ = v46
        }
        for fi47 := 0; fi47 < 10; fi47++ {
            v10 = v10
            _ = fi47
        }
    } else {
        v1 += (v1 % ((127) | 1))
    }
    switch 32767 {
    case 941:
        if true {
            v8 = append(v8, -827)
            var v48 int64 = 0
            _ = v48
        }
    default:
        var v49 string = ("\n\t" + "a")
        _ = v49
    }
    fmt.Println(v1, v1, v11)
    switch ("\n\t" + ("\n\t" + ("a" + ("hello" + (("\n\t" + ("foo bar" + "")) + "\n\t"))))) {
    case "foo bar":
        v50 := GenS3{F0: 6}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v44
    _ = v8
    _ = v9
}
