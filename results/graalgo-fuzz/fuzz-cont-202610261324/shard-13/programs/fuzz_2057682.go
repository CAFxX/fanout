package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "\n\t"
    v3 := []int{0, 32767}
    v4 := make(map[int]uint32)
    v5 := make(map[string]int)
    v2 = v2
    for fi6 := 0; fi6 < 10; fi6++ {
        v2 = "hello"
        v2 = (v2 + "foo bar")
        if (-32768 <= 0) {
            if false {
                var v7 float32 = 1.5
                if false {
                    fmt.Println(fi6, fi6, v2)
                }
                _ = v7
            }
            v2 = "\n\t"
        } else {
            var v8 uint = 255
            v8 += v8
            v9 := GenS1{F0: ("foo bar" + v2)}
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: v2}
    switch -1 {
    case 0:
        switch (v2 + "hello") {
        case "":
            for fi11 := 0; fi11 < 18; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    v2 = (v2 + v2)
                    _ = fi12
                }
                _ = fi11
            }
        case "foo bar":
            var v13 uint8 = 1
            _ = v13
        case "\n\t":
            fmt.Println(v2)
        default:
            _ = len(v1)
        }
    case 511:
        v2 = v2
    case 127:
        var v14 int64 = -9223372036854775808
        _ = v14
    default:
        v15 := GenS1{F0: ""}
        _ = v15
    }
    var v16 uint16 = 1
    v2 = v2
    if true {
        for fi17 := 0; fi17 < 1; fi17++ {
            switch (^(fi17)) {
            case -89:
                v2 = v2
            case 128:
                var v18 int16 = 32767
                _ = v18
            }
            v19 := GenS1{F0: v2}
            _ = fi17
            _ = v19
        }
        if false {
            v20 := GenS1{F0: v2}
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 uint64 = 1
                if len(v3) > 0 {
                    v3[0] = (fi21 % ((fi21) | 1))
                }
                _ = fi21
                _ = v22
            }
            _ = v20
        }
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 uint32 = 42
            var v25 string = ""
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    var v26 uint16 = (v16 << ((1) & 63))
    _ = len(v4)
    v2 = ((v2 + "") + v2)
    if false {
        fmt.Println(v2)
        if false {
            v16 -= 642
        } else {
            switch 684 {
            case -290:
                switch -881 {
                case 32768:
                    switch v2 {
                    case "a":
                        var v27 float32 = (0.0 - ((1.5 * 0.1) - (-2.25 * 1.5)))
                        _ = v27
                    case "hello":
                        switch v2 {
                        case "\n\t":
                            v3 = append(v3, 0)
                        }
                    case "":
                        var v28 bool = ((((1000 == 706) && (485 < 500)) || (42 < 2)) && (!(42 >= -86)))
                        _ = v28
                    default:
                        v16 = 787
                    }
                case 796:
                    var v29 uint64 = 18446744073709551615
                    _ = v29
                case 88:
                    v16 = (^(v26))
                default:
                    var v30 int8 = -82
                    _ = v30
                }
            case 221:
                v31 := GenS1{F0: v2}
                _ = v31
            default:
                v3 = append(v3, -684)
            }
            fmt.Println(v2)
            var v32 float32 = 1e10
            _ = v32
        }
        var v33 uint64 = 1
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
