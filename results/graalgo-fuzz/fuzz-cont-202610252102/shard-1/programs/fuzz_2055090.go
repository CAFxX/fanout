package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 int32; F1 uint8; F2 bool }
func genFn0() int64 {
    var v5 float32 = (float32(float64(3.14159)) - ((1.5 / -2.25) + float32(0.0)))
    if (!((3.14159 + 1e10) < 0.0)) {
        fmt.Println(v5)
    } else {
        fmt.Println(v5)
    }
    fmt.Println(v5)
    _ = v5
    return 905
}

func main() {
    var v1 uint8 = 1
    var v2 uint8 = 2
    var v3 int = 32768
    var v4 uint16 = 42
    v6 := []int{}
    if false {
        switch v3 {
        case -46:
            _ = len(v6)
        }
        v3 = (-(v3))
        v6 = append(v6, v3)
    }
    v7 := GenS1{F0: -2147483648, F1: -128, F2: -595}
    var v8 int16 = 191
    if len(v6) > 0 {
        v6[0] = (v3 - (-(v3)))
    }
    var v9 uint64 = 0
    var v10 uint64 = (v9 + v9)
    var v11 uint8 = 134
    var v12 uint8 = v1
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: int32(v3), F1: 91, F2: 0}
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: -639}
        v12 -= (v11 & 180)
        _ = fi13
        _ = v16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        if false {
            v6 = append(v6, fi17)
            v18 := GenS3{F0: 4294967295, F1: 2147483647, F2: "a"}
            if (42 < 1) {
                if false {
                    var v19 string = ""
                    for fi20 := 0; fi20 < 16; fi20++ {
                        var v21 int16 = 0
                        _ = fi20
                        _ = v21
                    }
                    _ = v19
                }
            } else {
                _ = len(v6)
                var v22 bool = false
                _ = v22
            }
            _ = v18
        } else {
            if (42 != 65535) {
                if true {
                    var v23 float64 = 1e10
                    _ = v23
                } else {
                    v12 = v1
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
        }
        _ = fi17
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v1 = 128
        _ = fi24
    }
    fmt.Println(v10, v4, v4)
    for fi25 := 0; fi25 < 9; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (false && ((42 == -128) || false)) {
                fmt.Println(fi25, v4)
                v3 = (-(fi26))
            } else {
                switch "hello" {
                case "":
                    if true {
                        var v27 int32 = 2147483647
                        v1 = (v2 - v2)
                        _ = v27
                    } else {
                        if (true && true) {
                            if (!(42 < fi25)) {
                                v28 := GenS4{F0: 42, F1: 42, F2: ((false && false) || true)}
                                v11 -= (v11 >> ((43) & 63))
                                fmt.Println(v12)
                                _ = v28
                            }
                            switch ("\n\t" + ("a" + ("a" + ("foo bar" + "")))) {
                            case "foo bar":
                                v29 := GenS2{F0: 740}
                                _ = v29
                            }
                        } else {
                            var v30 uint8 = (^(v2))
                            _ = v30
                        }
                        var v31 string = ("hello" + "")
                        _ = v31
                    }
                case "\n\t":
                    v12 -= v11
                case "foo bar":
                    var v32 int32 = 984
                    _ = v32
                }
            }
            v2 -= v2
            _ = fi26
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, fi25)
        _ = fi25
    }
    var v33 uint32 = 65535
    v12 = (v11 >> ((265) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
