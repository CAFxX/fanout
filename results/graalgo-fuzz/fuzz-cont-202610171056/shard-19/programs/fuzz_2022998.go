package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint64 {
    var v4 bool = ((0.1 * 1e10) >= 3.14159)
    if false {
        v5 := []int{1, 127, 127, 128}
        fmt.Println(v4)
        fmt.Println(v4)
        _ = v5
    } else {
        fmt.Println(v4)
        var v6 string = (((((("\n\t" + ((("foo bar" + "a") + ("" + ((("hello" + "a") + "foo bar") + (("\n\t" + "a") + "a")))) + "foo bar")) + "") + "\n\t") + "foo bar") + "a") + "foo bar")
        _ = v6
    }
    fmt.Println(v4)
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint = 90
    var v2 bool = (!(false && true))
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    fmt.Println(v1)
    v2 = ((((0.1 + ((3.14159 * 1e10) - float64(((1e10 * 0.1) * -2.25)))) + 0.1) + 1e10) != 0.0)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{int(v1), fi7}
        _ = fi7
        _ = v8
    }
    if ((0 <= 0) && v2) {
        v1 += 766
    } else {
        var v9 uint32 = 508
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 32767}
            var v13 int64 = 512
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(fi10, v2, v2)
        switch fi10 {
        case 32767:
            switch "\n\t" {
            case "":
                _ = len(v3)
            case "a":
                switch (-(fi10)) {
                case 2147483647:
                    switch "\n\t" {
                    case "foo bar":
                        var v14 string = "foo bar"
                        _ = v14
                    case "hello":
                        fmt.Println(v1, fi10)
                    }
                case 99:
                    for fi15 := 0; fi15 < 16; fi15++ {
                        var v16 string = "a"
                        var v17 string = v16
                        var v18 float32 = (0.1 - ((-2.25 / (0.1 - (3.14159 / -2.25))) - -2.25))
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                case 0:
                    v1 += v1
                default:
                    var v19 float64 = 0.0
                    _ = v19
                }
            case "\n\t":
                v2 = (42 != 32767)
            }
        case 451:
            if (!(!v2)) {
                v1 += v1
            }
        }
        _ = fi10
    }
    v2 = true
    var v20 uint32 = 103
    fmt.Println(v1, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v3
}
