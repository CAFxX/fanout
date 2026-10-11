package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0(p0 uint8, p1 string, p2 int8) uint16 {
    var v5 int8 = (p2 & 100)
    fmt.Println(p1)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint16 = 154
        v8 := GenS1{F0: v7, F1: 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 553
    var v3 bool = (-2147483648 != 956)
    v4 := make(map[uint32]int)
    var v9 uint8 = 2
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float32 = 1.5
            v13 := GenS1{F0: genFn0((v9 * v9), "hello", 0), F1: 2}
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int64 = 1
        _ = fi10
        _ = v14
    }
    var v15 float32 = -2.25
    var v16 bool = false
    if (192 < 0) {
        for fi17 := 0; fi17 < 14; fi17++ {
            switch fi17 {
            case 358:
                v16 = v3
            case 42:
                v18 := []int{fi17, fi17, 189, fi17}
                _ = v18
            case 9223372036854775807:
                if true {
                    switch fi17 {
                    case 2147483647:
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v20 := []int{}
                            _ = fi19
                            _ = v20
                        }
                    case -9223372036854775808:
                        var v21 string = "a"
                        _ = v21
                    default:
                        v22 := v4[620]
                        _ = v22
                    }
                } else {
                    v9 = 148
                }
            }
            _ = fi17
        }
        switch "hello" {
        case "\n\t":
            var v23 float32 = v15
            _ = v23
        case "foo bar":
            var v24 uint16 = 268
            _ = v24
        case "":
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint16 = 32767
                _ = fi25
                _ = v26
            }
        }
        v3 = false
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS1{F0: 65535, F1: 509}
        for fi29 := 0; fi29 < 4; fi29++ {
            if v3 {
                switch ("hello" + "hello") {
                case "hello":
                    var v30 string = "\n\t"
                    _ = v30
                case "foo bar":
                    var v31 int32 = 0
                    _ = v31
                }
            } else {
                v32 := []int{(fi27 / (((fi27 & (fi29 % ((fi27) | 1)))) | 1)), fi27}
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
                v15 -= ((v15 * 0.0) - float32(3.14159))
                v3 = ("hello" > "foo bar")
                fmt.Println(fi29, fi29)
                _ = v32
            }
            _ = fi29
        }
        var v33 int = (-(fi27))
        _ = fi27
        _ = v28
        _ = v33
    }
    v34 := GenS1{F0: uint16(v9), F1: 2}
    for fi35 := 0; fi35 < 16; fi35++ {
        fmt.Println(fi35, fi35, v15)
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 int16 = (^(v2))
            v38 := GenS1{F0: 0, F1: 255}
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    v9 -= v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v9
}
