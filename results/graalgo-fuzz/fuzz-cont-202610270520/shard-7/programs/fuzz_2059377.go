package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.1
    var v3 float64 = 1.5
    switch "hello" {
    case "":
        for fi4 := 0; fi4 < 13; fi4++ {
            fmt.Println(fi4, v2, fi4)
            _ = fi4
        }
    case "foo bar":
        if true {
            v2 = 1e10
            var v5 uint32 = 250
            _ = v5
        } else {
            fmt.Println(v2, v3)
            v6 := GenS1{F0: 2147483648}
            _ = v6
        }
    case "\n\t":
        v7 := GenS1{F0: 4294967295}
        _ = v7
    default:
        if (((("hello" + "") + ("" + ("\n\t" + ((((((("" + "a") + "hello") + ("a" + "foo bar")) + "") + "\n\t") + (("hello" + ("hello" + "")) + (("" + "foo bar") + (("hello" + (("foo bar" + "") + "a")) + "hello")))) + "")))) > "foo bar") && false) {
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 2; fi9++ {
                    v10 := []int{440, fi9, (-(fi9))}
                    _ = fi9
                    _ = v10
                }
                v2 += (1e10 * v2)
                _ = fi8
            }
            var v11 string = ""
            _ = v11
        } else {
            fmt.Println(v3, v2)
        }
    }
    v2 += v2
    var v12 uint = 42
    v13 := GenS1{F0: 4294967295}
    _ = len(v1)
    delete(v1, "a")
    v3 += v3
    v3 -= v3
    v3 += 0.1
    v2 = v2
    fmt.Println(v2, v2, v3)
    switch 32768 {
    case 127:
        v14 := GenS1{F0: 338}
        _ = v14
    case -854:
        var v15 uint64 = 0
        _ = v15
    case -444:
        var v16 int = int(v12)
        _ = v16
    }
    v12 += v12
    v2 -= float32((((v3 * -2.25) * (v3 / v3)) / (1.5 / v3)))
    v2 = v2
    if ((^(v12)) <= (v12 / (((^(v12))) | 1))) {
        if true {
            delete(v1, "a")
            if true {
                var v17 int = 4294967295
                v18 := []int{v17, (-(v17)), 88}
                _ = v17
                _ = v18
            }
        }
        fmt.Println(v12)
    }
    v19 := GenS1{F0: 804}
    if ((((("foo bar" + "hello") + "") + "foo bar") + "a") < "a") {
        var v20 int = -493
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
}
