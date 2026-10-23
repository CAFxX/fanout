package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ""
    var v3 int64 = 0
    var v4 uint8 = 79
    if ((1 != 4294967295) && true) {
        for fi5 := 0; fi5 < 10; fi5++ {
            if (1.5 > 1e10) {
                fmt.Println(v2, v4, fi5)
                v6 := v1[4294967295]
                _ = v6
            }
            _ = fi5
        }
    }
    v3 -= v3
    v4 -= (v4 | 127)
    var v7 uint = 923
    v2 = v2
    if false {
        switch 127 {
        case 32767:
            v7 = v7
        case 0:
            switch 4294967295 {
            case 546:
                v8 := []int{128}
                _ = v8
            case 996:
                var v9 float64 = 1e10
                _ = v9
            default:
                switch v2 {
                case "hello":
                    v10 := []int{32767, -681}
                    _ = v10
                case "a":
                    switch -9223372036854775808 {
                    case 0:
                        fmt.Println(v7, v4, v4)
                    case 128:
                        v1[42] = false
                    }
                }
            }
        }
        v11 := GenS1{F0: (2 < 255), F1: 40, F2: -122}
        fmt.Println(v2, v2, v2)
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        fmt.Println(v4)
        v13 := GenS1{F0: (993 >= 42), F1: 65535, F2: -15}
        _ = fi12
        _ = v13
    }
    if true {
        v3 += (v3 / ((v3) | 1))
        v2 = (("foo bar" + (v2 + "\n\t")) + v2)
        fmt.Println(v3, v7)
    }
    fmt.Println(v7, v4, v4)
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v2, v3)
        v2 = v2
        _ = fi14
    }
    v15 := GenS1{F0: (-505 > -122), F1: 2, F2: 42}
    fmt.Println(v4, v7, v7)
    v16 := GenS1{F0: false, F1: 42, F2: -5}
    v17 := GenS2{F0: (376 > 0), F1: v3}
    for fi18 := 0; fi18 < 13; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: true, F1: 2147483648, F2: 100}
            if (!(("foo bar" + v2) != v2)) {
                v4 = v4
                var v21 int64 = -1
                _ = v21
            } else {
                v1[2] = (!false)
                var v22 int32 = -966
                v7 += v7
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        var v23 string = ("foo bar" + "a")
        _ = fi18
        _ = v23
    }
    switch v2 {
    case "hello":
        for fi24 := 0; fi24 < 5; fi24++ {
            if true {
                v25 := GenS1{F0: true, F1: 2147483647, F2: 0}
                fmt.Println(v4, fi24, v4)
                _ = v25
            }
            var v26 int16 = -32768
            if (!(214 <= 0)) {
                v4 = 0
                v27 := GenS2{F0: true, F1: v3}
                _ = v27
            } else {
                v7 = v7
                v26 += v26
                if (908 < 0) {
                    var v28 float32 = 1.5
                    fmt.Println(fi24, v7)
                    _ = v28
                } else {
                    v29 := []int{int(v4), (fi24 >> ((672) & 63)), -1}
                    v30 := GenS2{F0: true, F1: (v3 * v3)}
                    _ = v29
                    _ = v30
                }
            }
            _ = fi24
            _ = v26
        }
    }
    v31 := GenS1{F0: (2147483648 != 444), F1: 145, F2: 59}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v7
}
