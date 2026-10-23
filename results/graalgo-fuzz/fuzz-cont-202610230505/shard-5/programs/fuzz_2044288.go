package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool }
func genFn0(p0 int64, p1 int8) uint64 {
    switch 127 {
    case 4294967295:
        v5 := GenS1{F0: 4294967295}
        _ = v5
    case 496:
        if (32767 != -694) {
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p1, p0)
            }
            if false {
                if ((false || (127 > 255)) || false) {
                    var v6 float32 = float32(-2.25)
                    _ = v6
                }
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 string = "\n\t"
                    _ = fi7
                    _ = v8
                }
                var v9 uint8 = 0
                _ = v9
            }
            p0 = (p0 + p0)
        } else {
            v10 := GenS1{F0: 255}
            var v11 uint = 0
            p0 = (p0 << ((2) & 63))
            _ = v10
            _ = v11
        }
    }
    var v12 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint64 = (v12 >> ((9) & 63))
        fmt.Println(v12, v12, p1)
        var v15 int = fi13
        _ = fi13
        _ = v14
        _ = v15
    }
    switch "a" {
    case "hello":
        v12 = 2
    }
    _ = p0
    _ = p1
    _ = v12
    return (v12 << ((18446744073709551615) & 63))
}
func genFn1() uint8 {
    v16 := GenS2{F0: -126}
    var v17 bool = (!true)
    var v18 int64 = -748
    fmt.Println(v18, v18)
    _ = v16
    _ = v17
    _ = v18
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = 0.1
    var v3 int = 29
    var v4 uint = 65535
    v4 = (^(v4))
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 string = ("" + ("foo bar" + ("hello" + ("foo bar" + "a"))))
        for fi21 := 0; fi21 < 14; fi21++ {
            fmt.Println(v2, fi19, v2)
            var v22 float32 = (-2.25 / (0.0 - (-2.25 + 1e10)))
            _ = fi21
            _ = v22
        }
        v3 = (-(v3))
        _ = fi19
        _ = v20
    }
    v23 := GenS1{F0: 4294967295}
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint32 = 962
        v26 := GenS3{F0: genFn1()}
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = 57
    v28 := GenS3{F0: genFn1()}
    v29 := []int{v3}
    v30 := GenS3{F0: 0}
    for fi31 := 0; fi31 < 8; fi31++ {
        switch ("hello" + "a") {
        case "hello":
            _ = len(v29)
        }
        var v32 uint = (^(v4))
        for fi33 := 0; fi33 < 8; fi33++ {
            var v34 int64 = -882
            if true {
                switch fi33 {
                case 9223372036854775807:
                    var v35 uint64 = genFn0(9223372036854775807, (v27 * (v27 >> (((v32 + 2)) & 63))))
                    _ = v35
                case -1:
                    var v36 uint32 = 4294967295
                    _ = v36
                default:
                    v37 := GenS4{F0: (1 >= 32767)}
                    _ = v37
                }
                v4 = (v4 / ((v32) | 1))
                switch "" {
                case "foo bar":
                    fmt.Println(v27, v2)
                default:
                    v38 := GenS3{F0: 2}
                    _ = v38
                }
                v2 -= 0.0
                fmt.Println(fi33, v2, v3)
            }
            v39 := GenS1{F0: v32}
            _ = fi33
            _ = v34
            _ = v39
        }
        _ = fi31
        _ = v32
    }
    var v40 uint16 = 656
    v2 += 1e10
    v29 = append(v29, (v3 & v3))
    if (239 == (v4 + 255)) {
        switch (^(v3)) {
        case 42:
            var v41 int32 = 2147483647
            _ = v41
        default:
            if true {
                v42 := GenS1{F0: (v4 & (^(v4)))}
                for fi43 := 0; fi43 < 8; fi43++ {
                    v44 := GenS2{F0: v27}
                    if false {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            switch 4294967295 {
                            case 32767:
                                var v46 uint = v4
                                _ = v46
                            case 127:
                                v47 := GenS3{F0: 42}
                                _ = v47
                            case -1:
                                fmt.Println(fi45)
                            }
                            _ = len(v29)
                            _ = fi45
                        }
                    }
                    _ = fi43
                    _ = v44
                }
                _ = v42
            }
        }
        v48 := GenS4{F0: (370 > 327)}
        _ = v48
    }
    var v49 float32 = (1.5 + 3.14159)
    for fi50 := 0; fi50 < 3; fi50++ {
        var v51 uint16 = 602
        fmt.Println(v2, v27, fi50)
        v27 -= (v27 << ((551) & 63))
        _ = fi50
        _ = v51
    }
    v52 := GenS4{F0: true}
    v40 -= (v40 % ((uint16(v4)) | 1))
    for fi53 := 0; fi53 < 11; fi53++ {
        _ = len(v29)
        _ = fi53
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v49
    _ = v52
}
