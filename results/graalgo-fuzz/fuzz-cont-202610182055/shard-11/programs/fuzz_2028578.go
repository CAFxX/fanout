package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint16 }
func genFn0() int32 {
    if (-513 == 580) {
        v5 := []int{4294967295, 42}
        _ = v5
    } else {
        v6 := []int{589, -453}
        v7 := GenS1{F0: 127, F1: 1, F2: (!false)}
        _ = v6
        _ = v7
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 bool = (9223372036854775807 != -893)
        _ = fi9
        _ = v10
    }
    _ = v8
    return 0
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 string = (("foo bar" + ("hello" + "hello")) + ("foo bar" + "hello"))
    var v4 float64 = 0.1
    for fi11 := 0; fi11 < 19; fi11++ {
        v4 = (0.0 + float64(float32(float64(0.0))))
        var v12 int32 = -997
        _ = fi11
        _ = v12
    }
    var v13 float64 = float64(3.14159)
    v3 = v3
    v4 = v4
    var v14 float64 = v4
    var v15 int64 = 893
    delete(v2, 42)
    v16 := GenS1{F0: 39, F1: genFn0(), F2: (0 < 1)}
    if false {
        v14 = 0.0
    }
    v4 = (0.0 + float64(0.1))
    if true {
        fmt.Println(v1, v13)
        for fi17 := 0; fi17 < 2; fi17++ {
            switch 990 {
            case 862:
                for fi18 := 0; fi18 < 9; fi18++ {
                    v19 := GenS1{F0: 0, F1: 2147483647, F2: v1}
                    var v20 int16 = 42
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            case -244:
                switch -1 {
                case 42:
                    for fi21 := 0; fi21 < 12; fi21++ {
                        switch 0 {
                        case 42:
                            fmt.Println(v14, fi21)
                        case 32768:
                            switch ((v3 + v3) + v3) {
                            case "a":
                                v22 := GenS1{F0: 2, F1: 2147483647, F2: (v3 <= v3)}
                                _ = v22
                            case "\n\t":
                                v4 += 0.1
                            case "hello":
                                var v23 int = -2147483648
                                _ = v23
                            }
                        }
                        v24 := GenS1{F0: 0, F1: 637, F2: (!(true && ((-890 == 0) || v1)))}
                        _ = fi21
                        _ = v24
                    }
                default:
                    var v25 float64 = v14
                    _ = v25
                }
            case 42:
                v2[0] = 4294967295
            }
            var v26 int16 = 32767
            _ = fi17
            _ = v26
        }
    } else {
        var v27 uint = 255
        v28 := []int{-2147483648}
        _ = v27
        _ = v28
    }
    if v1 {
        v29 := []int{9223372036854775807, -966}
        _ = v29
    }
    v30 := GenS2{F0: 67, F1: -31}
    v31 := GenS3{F0: 32767}
    v1 = ((v15 > v15) || v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
