package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 0
    v3 := []uint32{2147483648}
    var v4 bool = true
    v5 := make(map[string]uint32)
    var v6 float64 = 0.1
    v5[""] = 2
    var v7 string = "foo bar"
    fmt.Println(v4)
    v8 := v5[(("" + (("a" + "a") + "a")) + v7)]
    switch 32767 {
    case 559:
        v9 := GenS1{F0: -11, F1: 1}
        _ = v9
    case -210:
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS2{F0: uint32(v2), F1: 106, F2: v1}
            fmt.Println(v8, v7)
            _ = fi10
            _ = v11
        }
    case -1:
        switch 42 {
        case 32767:
            if false {
                delete(v5, "hello")
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch v7 {
                case "hello":
                    for fi12 := 0; fi12 < 3; fi12++ {
                        var v13 string = v7
                        v14 := v5[v13]
                        var v15 uint32 = v8
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                }
            } else {
                var v16 uint = 703
                _ = v16
            }
        case 32768:
            for fi17 := 0; fi17 < 16; fi17++ {
                if v4 {
                    v18 := GenS2{F0: v8, F1: 0, F2: (^(v1))}
                    v19 := GenS3{F0: 1}
                    _ = v18
                    _ = v19
                } else {
                    v6 = ((v6 * ((v6 / (float64(0.0) / (v6 - 1.5))) / 3.14159)) / 1.5)
                    delete(v5, "")
                }
                _ = fi17
            }
        case 42:
            v20 := v5[v7]
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 4; fi21++ {
        v6 = v6
        _ = fi21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        v2 += (v2 - (v2 >> ((1) & 63)))
        _ = fi22
    }
    switch (v7 + v7) {
    case "hello":
        var v23 int8 = 42
        _ = v23
    default:
        v6 -= ((v6 - (3.14159 + float64(1e10))) + v6)
    }
    v24 := v5[("" + v7)]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
