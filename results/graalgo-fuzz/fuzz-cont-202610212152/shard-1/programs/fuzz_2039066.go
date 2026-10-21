package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int; F2 uint }
func genFn0(p0 int64, p1 uint32, p2 bool) int16 {
    var v5 int32 = -141
    if p2 {
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: 54}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int32 = (v5 << ((1) & 63))
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -32768
}

func main() {
    v1 := []string{}
    var v2 int = -329
    v3 := []string{"\n\t", "", "\n\t"}
    var v4 int64 = -1
    for fi10 := 0; fi10 < 11; fi10++ {
        switch ("a" + "a") {
        case "hello":
            fmt.Println(fi10)
        case "\n\t":
            var v11 int = (fi10 | (fi10 << ((1) & 63)))
            _ = v11
        }
        var v12 uint = 65535
        for fi13 := 0; fi13 < 12; fi13++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v14 uint32 = 2147483648
            var v15 string = "hello"
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v12
    }
    var v16 uint8 = 127
    switch (^(v2)) {
    case 1:
        for fi17 := 0; fi17 < 15; fi17++ {
            v16 = (v16 ^ v16)
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(v4, v2, fi18)
                _ = fi18
            }
            v19 := GenS1{F0: -128}
            _ = fi17
            _ = v19
        }
    case 9223372036854775807:
        var v20 int16 = 0
        _ = v20
    default:
        v1 = append(v1, "hello")
    }
    v16 -= 247
    v16 -= 117
    var v21 uint8 = 181
    var v22 int = v2
    for fi23 := 0; fi23 < 12; fi23++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v24 int = v2
        _ = fi23
        _ = v24
    }
    switch "foo bar" {
    case "a":
        v25 := GenS1{F0: 127}
        _ = v25
    case "":
        if false {
            if (v4 >= 953) {
                v16 = v16
                var v26 float64 = (1e10 + 1e10)
                _ = v26
            } else {
                var v27 uint8 = (^(v16))
                var v28 bool = false
                for fi29 := 0; fi29 < 7; fi29++ {
                    v30 := GenS1{F0: int8(v22)}
                    _ = fi29
                    _ = v30
                }
                _ = v27
                _ = v28
            }
            v31 := GenS1{F0: -1}
            for fi32 := 0; fi32 < 1; fi32++ {
                v22 = v22
                fmt.Println(fi32, v4)
                _ = fi32
            }
            _ = v31
        } else {
            v33 := GenS2{F0: 984}
            switch v22 {
            case 9223372036854775807:
                v21 = v21
            case -1:
                if true {
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v2 = v22
                        fmt.Println(v4)
                        _ = fi34
                    }
                    for fi35 := 0; fi35 < 15; fi35++ {
                        if len(v3) > 0 {
                            v3[0] = (("\n\t" + "a") + "hello")
                        }
                        _ = fi35
                    }
                    var v36 int = (v22 ^ -1)
                    _ = v36
                } else {
                    var v37 bool = (949 != 2)
                    _ = v37
                }
            }
            v38 := GenS1{F0: -94}
            _ = v33
            _ = v38
        }
    default:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    var v39 float64 = 0.0
    if true {
        if (-335 > -62) {
            switch (-(v22)) {
            case -912:
                for fi40 := 0; fi40 < 7; fi40++ {
                    var v41 uint8 = v21
                    _ = fi40
                    _ = v41
                }
            case 323:
                var v42 bool = (false && ((false && true) && false))
                _ = v42
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v43 float32 = (((3.14159 + 0.1) / 1.0) / float32(float64(3.14159)))
            v44 := GenS3{F0: "a", F1: false, F2: -2147483648}
            _ = v43
            _ = v44
        }
        v22 -= 0
        v45 := GenS4{F0: int32(v2)}
        _ = v45
    } else {
        v39 -= 1e10
        v21 += 127
        for fi46 := 0; fi46 < 9; fi46++ {
            fmt.Println(v4, v2, v22)
            if (42 >= 18446744073709551615) {
                if (-1 == 42) {
                    for fi47 := 0; fi47 < 4; fi47++ {
                        if len(v3) > 0 {
                            v3[0] = ("foo bar" + "a")
                        }
                        _ = fi47
                    }
                } else {
                    var v48 uint32 = 0
                    v49 := GenS5{F0: "\n\t", F1: v22, F2: 459}
                    v50 := GenS4{F0: -2147483648}
                    _ = v48
                    _ = v49
                    _ = v50
                }
                var v51 bool = (!(!false))
                v16 += (^(v21))
                _ = v51
            } else {
                v21 = (v16 | (v21 << ((65535) & 63)))
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                _ = len(v3)
            }
            _ = fi46
        }
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        var v53 int8 = 42
        var v54 int32 = 42
        v4 = (v4 * -464)
        _ = fi52
        _ = v53
        _ = v54
    }
    for fi55 := 0; fi55 < 14; fi55++ {
        v2 -= v22
        if true {
            fmt.Println(v2, v4, v16)
            var v56 int32 = -529
            _ = v56
        } else {
            _ = len(v3)
        }
        _ = fi55
    }
    v57 := GenS3{F0: "a", F1: true, F2: -462}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v57
}
