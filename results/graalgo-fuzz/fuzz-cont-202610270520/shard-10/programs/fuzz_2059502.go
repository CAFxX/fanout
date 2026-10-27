package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int {
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = 52
        v7 := GenS1{F0: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch ("hello" + "\n\t") {
    case "a":
        var v8 float64 = 3.14159
        _ = v8
    case "hello":
        if true {
            v9 := GenS1{F0: 253}
            v10 := GenS1{F0: 0}
            _ = v10
            _ = v9
        }
    default:
        v11 := []int{42, 782, 957}
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := GenS1{F0: -554}
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = 63
                _ = fi14
                _ = v15
            }
            var v16 uint16 = 32767
            _ = fi12
            _ = v13
            _ = v16
        }
        var v17 int16 = 0
        if ((0.1 / (1e10 / 1.0)) <= 1e10) {
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := GenS1{F0: -1}
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 uint64 = 2
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
        }
        _ = v17
    } else {
        var v22 bool = false
        _ = v22
    }
    return 42
}

func main() {
    v1 := []string{"foo bar", ("" + (("a" + "\n\t") + "")), "\n\t", "foo bar"}
    var v2 uint32 = 163
    var v3 int16 = 1000
    v4 := make(map[int]uint32)
    fmt.Println(v2, v3)
    if (true || (-2147483648 >= 2147483647)) {
        fmt.Println(v3)
        v3 = (v3 ^ v3)
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                v25 := GenS1{F0: 1}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    }
    if (!(116 > 128)) {
        fmt.Println(v2, v3)
    } else {
        v3 += (v3 - v3)
        v3 += (v3 ^ v3)
    }
    switch "\n\t" {
    case "hello":
        v2 = v2
    case "":
        v2 = 2147483648
    case "\n\t":
        v2 = (v2 >> ((407) & 63))
    }
    delete(v4, genFn0())
    v3 = 0
    var v26 bool = (2 <= 1)
    fmt.Println(v3, v3)
    _ = len(v4)
    v27 := GenS1{F0: -582}
    if (2 <= 102) {
        for fi28 := 0; fi28 < 2; fi28++ {
            if (1 < 1) {
                if (!v26) {
                    var v29 float32 = 1.5
                    fmt.Println(v26, v26, v2)
                    _ = v29
                }
            } else {
                fmt.Println(v26, v2)
                fmt.Println(v26)
            }
            v30 := GenS1{F0: -16}
            _ = fi28
            _ = v30
        }
    }
    v26 = v26
    switch "foo bar" {
    case "":
        v2 = (^(v2))
    case "foo bar":
        var v31 float64 = 1e10
        _ = v31
    case "a":
        if v26 {
            v2 = v2
            switch ("a" + "\n\t") {
            case "hello":
                fmt.Println(v3, v3)
            case "":
                var v32 float64 = 0.0
                _ = v32
            case "\n\t":
                var v33 uint = 65535
                _ = v33
            }
        } else {
            var v34 int16 = (v3 / ((1) | 1))
            v3 = v3
            var v35 uint16 = 1
            _ = v34
            _ = v35
        }
    default:
        switch -2147483648 {
        case 2147483647:
            for fi36 := 0; fi36 < 12; fi36++ {
                var v37 uint = 65535
                v2 -= v2
                _ = fi36
                _ = v37
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
