package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 bool) uint16 {
    v7 := GenS1{F0: 2147483647, F1: 4294967295}
    switch "foo bar" {
    case "\n\t":
        p0 = (0.1 <= 0.0)
    case "foo bar":
        switch 32767 {
        case 32767:
            p0 = (223 >= 1)
        case 127:
            switch "\n\t" {
            case "a":
                v8 := GenS2{F0: (0 <= 918), F1: 459}
                _ = v8
            }
        }
    case "a":
        p0 = (("hello" + ((("a" + "\n\t") + "foo bar") + ("\n\t" + "foo bar"))) > "hello")
    }
    p0 = true
    for fi9 := 0; fi9 < 13; fi9++ {
        p0 = p0
        _ = fi9
    }
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 2
}

func main() {
    var v1 string = ""
    var v2 int32 = -2147483648
    var v3 bool = false
    var v4 int32 = v2
    var v5 uint16 = 42
    var v6 uint = 1
    v10 := GenS3{F0: 593, F1: -5}
    v3 = v3
    if ((v1 <= v1) && (v3 && (v3 && v3))) {
        v6 = v6
        var v11 uint64 = 803
        if v3 {
            if (1 >= 1) {
                var v12 float32 = 0.0
                v13 := []int{1, 127, 2147483647, 9223372036854775807}
                _ = v12
                _ = v13
            }
            var v14 uint = v6
            switch -2147483648 {
            case 4294967295:
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v14, v6, fi15)
                    var v16 int64 = 9223372036854775807
                    var v17 int64 = (v16 << ((v14) & 63))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v14
        } else {
            v1 = "a"
            for fi18 := 0; fi18 < 11; fi18++ {
                if (v6 <= v6) {
                    var v19 uint32 = 2147483648
                    _ = v19
                }
                _ = fi18
            }
            v5 -= v5
        }
        _ = v11
    }
    fmt.Println(v2, v2)
    fmt.Println(v3, v2, v5)
    var v20 int16 = 1
    v2 = v2
    var v21 uint8 = 2
    v22 := GenS3{F0: 369, F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
