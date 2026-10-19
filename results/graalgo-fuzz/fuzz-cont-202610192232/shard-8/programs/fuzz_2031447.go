package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int16; F2 int32 }
func genFn0() uint8 {
    switch ("" + "a") {
    case "a":
        v4 := []int{1, -120, -719}
        _ = v4
    case "":
        var v5 string = "hello"
        _ = v5
    case "hello":
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := GenS1{F0: false, F1: 1, F2: -32768}
            if (-196 > 1) {
                if ("" >= "") {
                    v8 := []int{fi6, fi6, (fi6 % (((fi6 & (fi6 * 32768))) | 1)), -2147483648}
                    _ = v8
                }
                v9 := []int{}
                var v10 uint32 = 4294967295
                _ = v10
                _ = v9
            }
            var v11 int64 = -537
            _ = fi6
            _ = v11
            _ = v7
        }
    default:
        var v12 uint = 2
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint32 = 0
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint32 = (v14 - (v14 << ((0) & 63)))
            var v17 string = ""
            _ = fi15
            _ = v16
            _ = v17
        }
        v18 := []int{-942, -1, (fi13 >> ((255) & 63))}
        _ = fi13
        _ = v14
        _ = v18
    }
    var v19 uint = 255
    _ = v19
    return 1
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]bool)
    var v3 uint16 = 1
    var v20 uint64 = 2
    v3 = v3
    v20 -= v20
    if (false && (!(2 != 18446744073709551615))) {
        var v21 float32 = v1
        if false {
            var v22 uint16 = v3
            _ = v22
        } else {
            for fi23 := 0; fi23 < 3; fi23++ {
                v24 := GenS1{F0: (173 <= 255), F1: 18446744073709551615, F2: 0}
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 8; fi25++ {
                if false {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if (!((false && (true || (((!false) && (4294967295 > (v20 ^ v20))) || (!false)))) || false)) {
                            v27 := GenS2{F0: 2147483647, F1: 4294967295, F2: ("hello" + "hello")}
                            v20 += v20
                            _ = v27
                        } else {
                            var v28 int64 = 1
                            _ = v28
                        }
                        var v29 uint8 = 127
                        _ = fi26
                        _ = v29
                    }
                } else {
                    v2[(fi25 % ((fi25) | 1))] = false
                    var v30 uint32 = 2147483648
                    _ = v30
                }
                var v31 uint = uint(v20)
                _ = fi25
                _ = v31
            }
        }
        v20 = v20
        _ = v21
    } else {
        v32 := []int{127, -2147483648}
        var v33 int16 = 0
        if (("\n\t" + "a") == ("hello" + ("hello" + ("hello" + ("foo bar" + ""))))) {
            switch 833 {
            case -2147483648:
                v1 -= ((v1 - v1) + (((v1 * float32((3.14159 / (3.14159 + (-2.25 + (0.1 - (0.1 - (0.1 - -2.25)))))))) - v1) * v1))
            default:
                _ = len(v2)
            }
        }
        _ = v32
        _ = v33
    }
    switch "\n\t" {
    case "a":
        if false {
            switch -660 {
            case -2147483648:
                var v34 int64 = -1
                _ = v34
            case 127:
                fmt.Println(v20, v3)
            default:
                v3 = 32767
            }
            if ((!(4294967295 < 4294967295)) || false) {
                v20 += v20
                var v35 int16 = -32768
                _ = v35
            } else {
                v36 := GenS1{F0: (!true), F1: v20, F2: 1}
                if false {
                    var v37 uint8 = 30
                    fmt.Println(v37, v20, v3)
                    _ = v37
                }
                v38 := GenS3{F0: -595}
                _ = v36
                _ = v38
            }
        } else {
            v3 += (v3 + (^(v3)))
            for fi39 := 0; fi39 < 1; fi39++ {
                var v40 uint32 = 947
                _ = fi39
                _ = v40
            }
        }
    case "hello":
        v41 := GenS4{F0: 42, F1: -196, F2: -140}
        _ = v41
    case "":
        v1 = (float32((((0.0 * ((1.5 + 1.5) * 3.14159)) + (1.5 / 0.1)) * -2.25)) / (float32(3.14159) - 0.0))
    default:
        for fi42 := 0; fi42 < 1; fi42++ {
            v3 = 448
            _ = fi42
        }
    }
    v43 := []int{}
    v1 = float32(1.5)
    v1 = (3.14159 - v1)
    v44 := GenS2{F0: 2, F1: 1, F2: ""}
    if len(v43) > 0 {
        v43[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v43
    _ = v44
}
