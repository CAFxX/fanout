package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int, p1 bool, p2 uint64) int16 {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (fi7 | p0) {
        case -2147483648:
            fmt.Println(p1)
        }
        _ = fi7
    }
    if true {
        fmt.Println(p1, p1)
        v8 := []int{-595, 42, (p0 & -2147483648)}
        _ = v8
    }
    v9 := []int{p0, p0, (^(p0)), p0}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 uint32) uint64 {
    v10 := []int{}
    switch -9223372036854775808 {
    case 32768:
        if true {
            fmt.Println(p1, p1)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v11 uint8 = 42
            _ = v11
        } else {
            p1 = 4294967295
            for fi12 := 0; fi12 < 10; fi12++ {
                p0 = p0
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint8 = 247
                    var v15 uint16 = 65535
                    var v16 int = 704
                    var v17 int = v16
                    v18 := GenS1{F0: v15, F1: 20}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                p1 = p1
                _ = fi12
            }
        }
    case 9223372036854775807:
        p1 = p1
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v19 uint16 = 0
    var v20 uint = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v19
    _ = v20
    return (^(p0))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 42
    var v3 float64 = float64((((-2.25 + 3.14159) / 3.14159) * 3.14159))
    var v4 uint = 1
    v5 := make(map[uint32]int)
    var v6 float64 = v3
    for fi21 := 0; fi21 < 2; fi21++ {
        v22 := GenS2{F0: -1, F1: false, F2: 253}
        if (!true) {
            var v23 int = fi21
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v6 = v3
    switch "a" {
    case "":
        switch "" {
        case "":
            v6 -= v3
        }
    case "a":
        var v24 uint = 4294967295
        _ = v24
    case "hello":
        var v25 uint64 = 544
        _ = v25
    }
    v26 := GenS2{F0: 21, F1: (633 == -1), F2: 0}
    switch -699 {
    case 114:
        if false {
            v4 = v4
        }
    case 32768:
        var v27 float64 = (v6 + v6)
        _ = v27
    default:
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 uint8 = (v2 * v2)
            v30 := GenS2{F0: 56, F1: true, F2: 42}
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    if ((42 <= -329) || ((false || (!false)) || true)) {
        v4 = (v4 | (v4 ^ 255))
        v31 := GenS3{F0: 42}
        _ = v31
    } else {
        if true {
            delete(v5, 569)
            var v32 int8 = 10
            if true {
                var v33 string = "a"
                v34 := GenS2{F0: (v32 + (v32 * 127)), F1: true, F2: 42}
                switch (v33 + v33) {
                case "a":
                    v35 := []int{1, -2147483648, 551, 2147483647}
                    _ = v35
                }
                _ = v33
                _ = v34
            } else {
                if (((("hello" + "") + (("foo bar" + "\n\t") + ("hello" + ("" + "a")))) + "hello") > "foo bar") {
                    fmt.Println(v4, v3, v2)
                }
                switch 950 {
                case 32767:
                    var v36 uint64 = 4294967295
                    _ = v36
                case -2147483648:
                    if (!false) {
                        var v37 int8 = 100
                        v38 := []int{int(v4), 32768}
                        _ = v37
                        _ = v38
                    } else {
                        var v39 uint = 0
                        v40 := []int{}
                        _ = len(v5)
                        _ = v39
                        _ = v40
                    }
                case 128:
                    var v41 uint = v4
                    _ = v41
                default:
                    var v42 uint32 = 4294967295
                    _ = v42
                }
            }
            _ = v32
        }
        v6 += 0.0
        v43 := GenS3{F0: genFn0(32767, (1 >= 20), 42)}
        _ = v43
    }
    var v44 int64 = -1
    fmt.Println(v3)
    v45 := []int{32767, 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
