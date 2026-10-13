package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 string, p2 uint16) uint8 {
    var v7 uint8 = 255
    switch p1 {
    case "hello":
        var v8 uint16 = (p2 % (((p2 / ((p2) | 1))) | 1))
        _ = v8
    case "\n\t":
        switch 127 {
        case -1:
            var v9 int8 = (p0 | int8(v7))
            _ = v9
        case 777:
            var v10 float32 = 0.0
            _ = v10
        default:
            switch "a" {
            case "":
                var v11 int32 = -1
                _ = v11
            case "\n\t":
                for fi12 := 0; fi12 < 13; fi12++ {
                    fmt.Println(p1, p1)
                    p2 = 0
                    _ = fi12
                }
            }
        }
    case "foo bar":
        fmt.Println(p0, p0, p2)
    }
    var v13 uint = 255
    v14 := []int{9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return (^(v7))
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    var v4 int64 = 361
    var v5 string = ("\n\t" + ("\n\t" + (("hello" + ("a" + ("hello" + "hello"))) + "")))
    var v6 float64 = 0.1
    delete(v3, 774)
    var v15 int8 = 0
    v16 := GenS1{F0: (^(v15))}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 uint16 = 788
        _ = len(v1)
        _ = fi17
        _ = v18
    }
    if (65535 == 42) {
        for fi19 := 0; fi19 < 13; fi19++ {
            v6 = v6
            if false {
                v5 = "hello"
                v20 := GenS1{F0: v15}
                _ = v20
            } else {
                var v21 int16 = 1000
                fmt.Println(v5)
                _ = v21
            }
            var v22 uint = 255
            _ = fi19
            _ = v22
        }
        if (42 == 128) {
            fmt.Println(v15, v2, v5)
        }
    }
    v23 := GenS2{F0: 42}
    v24 := GenS3{F0: 2}
    v6 -= 1.5
    if false {
        fmt.Println(v15, v6, v4)
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(v4, v4)
            v26 := GenS1{F0: (v15 | (v15 * (v15 & (v15 * v15))))}
            switch (^(fi25)) {
            case 4294967295:
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 string = "a"
                    var v29 uint = 73
                    var v30 int64 = v2
                    _ = fi27
                    _ = v28
                    _ = v29
                    _ = v30
                }
            }
            _ = fi25
            _ = v26
        }
        v15 = -1
    } else {
        var v31 int64 = -1
        switch 100 {
        case 128:
            switch -2147483648 {
            case 1:
                var v32 int64 = v31
                _ = v32
            case 492:
                if true {
                    fmt.Println(v6, v15)
                } else {
                    v3[-9223372036854775808] = 32768
                }
            case -9223372036854775808:
                v33 := GenS2{F0: 45}
                _ = v33
            }
        default:
            fmt.Println(v15, v2, v31)
        }
        _ = v31
    }
    var v34 float64 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
