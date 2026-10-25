package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint16, p1 uint, p2 uint8) uint {
    fmt.Println(p0)
    if true {
        fmt.Println(p2)
        v5 := []int{}
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        var v7 uint16 = 1
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := []int{42, 921}
    v4 := []string{"a", ("hello" + ("a" + "\n\t")), "a"}
    var v8 int32 = -532
    var v9 uint8 = 42
    switch (("" + "foo bar") + "\n\t") {
    case "\n\t":
        v10 := GenS1{F0: ((^(v9)) < v9), F1: 4294967295}
        _ = v10
    case "a":
        var v11 uint8 = 205
        _ = v11
    case "hello":
        var v12 int16 = -32768
        _ = v12
    default:
        var v13 int64 = 630
        _ = v13
    }
    var v14 string = ""
    v14 = v14
    v9 = (v9 * 1)
    for fi15 := 0; fi15 < 13; fi15++ {
        fmt.Println(v8, fi15)
        _ = fi15
    }
    if (true || true) {
        var v16 int64 = -927
        if false {
            var v17 float32 = (0.0 * 0.0)
            _ = v17
        } else {
            _ = len(v1)
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                if true {
                    v16 += v16
                }
                var v20 uint16 = 42
                _ = fi19
                _ = v20
            }
            v16 -= v16
            fmt.Println(v9, v9)
            _ = fi18
        }
        _ = v16
    }
    if (1e10 == 3.14159) {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 int16 = 42
            _ = fi21
            _ = v22
        }
        fmt.Println(v9)
    } else {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch (v14 + v14) {
            case "\n\t":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "foo bar":
                v9 = (v9 * (v9 >> ((2) & 63)))
            }
            v1[862] = (v14 + (v14 + v14))
            for fi24 := 0; fi24 < 12; fi24++ {
                if ("foo bar" < v14) {
                    v14 = "hello"
                }
                var v25 int = 721
                switch (v25 | 2147483647) {
                case 196:
                    v26 := GenS1{F0: (-2147483648 == (fi24 | fi23)), F1: 1}
                    _ = v26
                case 42:
                    v27 := v2[2]
                    _ = v27
                case 128:
                    fmt.Println(v14, v14)
                }
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    }
    v9 += v9
    for fi28 := 0; fi28 < 15; fi28++ {
        if false {
            v14 = (v14 + v14)
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            var v30 int16 = 1000
            fmt.Println(fi28, fi29)
            switch (fi29 ^ 970) {
            case 0:
                switch "" {
                case "":
                    if false {
                        v30 = (-(v30))
                        v14 = v14
                    }
                }
            case 9223372036854775807:
                var v31 bool = (200 <= 42)
                _ = v31
            }
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
