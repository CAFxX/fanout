package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
func genFn0() string {
    var v7 int64 = -9223372036854775808
    v8 := GenS1{F0: 0, F1: (!false)}
    var v9 uint8 = 255
    var v10 int16 = -450
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ""
}
func genFn1() uint32 {
    if (138 == 0) {
        var v11 uint = 415
        _ = v11
    }
    var v12 uint64 = 2
    for fi13 := 0; fi13 < 17; fi13++ {
        v12 -= v12
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int8 = -1
            v16 := GenS2{F0: (("" + "a") + "\n\t")}
            v15 = 1
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    v17 := GenS2{F0: "hello"}
    _ = v12
    _ = v17
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 string = (("" + "") + ("hello" + "foo bar"))
    var v5 uint = 18446744073709551615
    var v6 uint = (v5 & (v5 ^ (v5 << (((^(v5))) & 63))))
    switch -269 {
    case 0:
        var v18 uint16 = 78
        _ = v18
    default:
        v5 -= v6
    }
    v19 := []int{127, 32768}
    v6 += 255
    v6 += (v5 << ((v6) & 63))
    if true {
        for fi20 := 0; fi20 < 8; fi20++ {
            if (2 < 107) {
                fmt.Println(v4)
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                switch fi20 {
                case 127:
                    v5 -= 0
                case -9223372036854775808:
                    if (((1e10 + 0.1) * 0.0) != 1e10) {
                        var v21 uint8 = 42
                        var v22 string = "hello"
                        var v23 int = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        if false {
                            var v24 uint8 = 17
                            switch v4 {
                            case "foo bar":
                                fmt.Println(v5)
                            default:
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    var v26 uint16 = 42
                                    _ = fi25
                                    _ = v26
                                }
                            }
                            _ = v24
                        } else {
                            fmt.Println(v4, fi20, v5)
                            fmt.Println(v5, fi20)
                            var v27 string = ((v4 + ((("foo bar" + v4) + "a") + "")) + "hello")
                            _ = v27
                        }
                    }
                }
            }
            v28 := GenS3{F0: 0}
            v5 = v5
            _ = fi20
            _ = v28
        }
        var v29 bool = false
        _ = v29
    } else {
        var v30 bool = ((1 != 836) || ((32767 == 0) || (65535 < 0)))
        switch 4294967295 {
        case 32767:
            if len(v19) > 0 {
                v19[0] = 32767
            }
        }
        _ = v30
    }
    v19 = append(v19, 127)
    fmt.Println(v6, v4)
    for fi31 := 0; fi31 < 13; fi31++ {
        if true {
            var v32 float32 = 1e10
            _ = v32
        }
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi31
    }
    v5 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
