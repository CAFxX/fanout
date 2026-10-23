package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
func genFn0() bool {
    if (0 < 123) {
        var v5 int64 = 9223372036854775807
        var v6 int8 = -1
        _ = v5
        _ = v6
    } else {
        v7 := GenS1{F0: (1 < 127)}
        v8 := []int{-1}
        switch "a" {
        case "hello":
            var v9 int64 = -1
            _ = v9
        case "foo bar":
            var v10 uint = 881
            _ = v10
        default:
            if (18446744073709551615 == 1) {
                for fi11 := 0; fi11 < 8; fi11++ {
                    var v12 int = 614
                    fmt.Println(fi11)
                    _ = fi11
                    _ = v12
                }
                var v13 int = 1
                _ = v13
            } else {
                var v14 float32 = 1.5
                _ = v14
            }
        }
        _ = v7
        _ = v8
    }
    var v15 uint32 = 2
    if (int64(v15) <= 1) {
        v16 := GenS2{F0: 610}
        if (false || (2 == 195)) {
            fmt.Println(v15)
        } else {
            var v17 bool = (-2147483648 >= 42)
            _ = v17
        }
        _ = v16
    }
    _ = v15
    return (v15 <= (v15 & 565))
}

func main() {
    var v1 float32 = 0.0
    var v2 uint16 = 65535
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]bool)
    fmt.Println(v1, v1, v3)
    v18 := GenS2{F0: (v2 << ((42) & 63))}
    v19 := v4[(("\n\t" + ("hello" + ((("a" + ("\n\t" + (("hello" + "") + ("hello" + "\n\t")))) + "") + "hello"))) + ("\n\t" + "foo bar"))]
    var v20 float32 = v1
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS3{F0: 4294967295}
        _ = fi21
        _ = v22
    }
    switch 32767 {
    case -216:
        var v23 float64 = ((((-2.25 - 0.0) + ((float64(0.0) * float64(0.0)) * 1.5)) - ((0.0 - ((1.5 - (0.1 - float64(v1))) * 1.5)) - 1.5)) - ((0.0 / 1e10) / 1e10))
        _ = v23
    case 32768:
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v1, fi24, v2)
            var v25 int8 = 44
            _ = fi24
            _ = v25
        }
    default:
        v19 = (v19 || (-1 == -32768))
    }
    if (-471 < 1000) {
        var v26 uint8 = 128
        v2 -= (v2 << ((4294967295) & 63))
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            for fi28 := 0; fi28 < 8; fi28++ {
                v19 = v19
                v29 := []int{32767, (fi28 >> ((599) & 63)), fi28, 389}
                var v30 float32 = (1e10 - 1.5)
                _ = fi28
                _ = v29
                _ = v30
            }
            v31 := GenS1{F0: true}
            fmt.Println(v1, v1, v1)
            _ = fi27
            _ = v31
        }
        var v32 string = "a"
        _ = v32
    }
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
