package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint64) uint16 {
    var v5 string = (("\n\t" + "hello") + "\n\t")
    var v6 int = 1
    switch (v6 | int(p0)) {
    case 871:
        v7 := GenS1{F0: -802}
        _ = v7
    default:
        v8 := []int{v6, v6, int(p0)}
        _ = v8
    }
    p0 = 4294967295
    _ = p0
    _ = v5
    _ = v6
    return 316
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 int32 = 0
    v3 := make(map[string]string)
    var v4 int16 = -490
    var v9 uint32 = 464
    v1 -= v1
    var v10 uint16 = 65535
    var v11 float32 = (-2.25 + (3.14159 - float32(-2.25)))
    v2 -= v2
    if (((v4 * 1) >= (v4 ^ (v4 | v4))) || (false && (127 < 1))) {
        v12 := []int{0}
        var v13 int16 = (v4 | (v4 / ((0) | 1)))
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v9 = 4294967295
        _ = fi14
    }
    if false {
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                v2 += (v2 << ((255) & 63))
                _ = fi16
            }
            v11 -= float32(-2.25)
            _ = fi15
        }
        if (255 >= (v10 - uint16(v1))) {
            v2 = (v2 ^ (v2 - 2147483647))
            fmt.Println(v4, v10)
        } else {
            if true {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v18 := v3[("" + "foo bar")]
                    var v19 string = (v18 + ((v18 + "hello") + "a"))
                    v10 -= (^(v10))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint8 = 158
                v21 := []int{}
                _ = v20
                _ = v21
            } else {
                switch 32767 {
                case 671:
                    if true {
                        for fi22 := 0; fi22 < 5; fi22++ {
                            var v23 int = fi22
                            if false {
                                v24 := []int{}
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                        var v25 int = int(v1)
                        if (v4 >= 42) {
                            v2 += (v2 >> ((v1) & 63))
                            if false {
                                var v26 uint32 = v9
                                _ = v26
                            }
                            v27 := GenS1{F0: 1}
                            _ = v27
                        } else {
                            _ = len(v3)
                        }
                        _ = v25
                    } else {
                        fmt.Println(v4)
                        var v28 uint64 = 580
                        _ = v28
                    }
                case -773:
                    v4 += int16(v10)
                default:
                    switch "hello" {
                    case "":
                        var v29 float64 = 0.1
                        _ = v29
                    default:
                        var v30 bool = false
                        _ = v30
                    }
                }
                var v31 int32 = (v2 - (v2 << ((v1) & 63)))
                _ = v31
            }
            fmt.Println(v2, v9)
        }
    } else {
        var v32 bool = false
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        fmt.Println(v1, v11)
        v4 = (v4 * v4)
        v10 = v10
        _ = fi33
    }
    if false {
        v4 -= v4
        for fi34 := 0; fi34 < 7; fi34++ {
            if false {
                fmt.Println(v11, fi34)
                if false {
                    var v35 string = ""
                    fmt.Println(fi34)
                    _ = v35
                } else {
                    v4 -= -948
                    fmt.Println(v10)
                }
                if (42 == 2) {
                    var v36 float64 = 0.1
                    _ = v36
                }
            } else {
                var v37 int8 = 68
                var v38 int64 = int64(fi34)
                _ = v37
                _ = v38
            }
            if false {
                fmt.Println(v2, v10)
                var v39 int16 = -822
                _ = v39
            } else {
                var v40 int16 = 32767
                _ = v40
            }
            v41 := GenS1{F0: 217}
            _ = fi34
            _ = v41
        }
        var v42 uint16 = 65535
        _ = v42
    }
    if true {
        v43 := GenS2{F0: 236, F1: 223, F2: 1}
        _ = len(v3)
        _ = v43
    } else {
        v2 -= -2147483648
    }
    fmt.Println(v9, v4)
    _ = len(v3)
    if true {
        var v44 int32 = v2
        v9 = (v9 << ((v1) & 63))
        _ = v44
    }
    switch "hello" {
    case "foo bar":
        var v45 uint = (v1 % ((1) | 1))
        _ = v45
    case "a":
        delete(v3, ("\n\t" + (("foo bar" + "hello") + (("" + "hello") + ""))))
    }
    var v46 int16 = (v4 & -646)
    for fi47 := 0; fi47 < 2; fi47++ {
        v48 := GenS3{F0: int8(fi47)}
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v9
}
