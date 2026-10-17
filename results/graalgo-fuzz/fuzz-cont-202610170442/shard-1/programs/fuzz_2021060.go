package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 string, p1 bool) uint32 {
    if false {
        var v4 float32 = float32(-2.25)
        _ = v4
    } else {
        var v5 uint = 255
        _ = v5
    }
    var v6 int64 = 42
    v7 := GenS1{F0: (p0 + p0)}
    v8 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 127
    var v3 int32 = -12
    var v9 uint16 = 1
    switch 32768 {
    case 32767:
        v10 := GenS2{F0: -32768}
        _ = v10
    case -1:
        v3 += (v3 ^ (v3 % ((42) | 1)))
    case 363:
        switch 4294967295 {
        case 701:
            v11 := GenS2{F0: -76}
            _ = v11
        case 32767:
            v2 = (^(v2))
        }
    }
    if (false || false) {
        v9 -= 1
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float32 = 0.1
        v14 := GenS1{F0: ("" + "")}
        _ = fi12
        _ = v13
        _ = v14
    }
    if (100 == 42) {
        for fi15 := 0; fi15 < 12; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int16 = int16(fi15)
                _ = fi16
                _ = v17
            }
            v1["foo bar"] = "\n\t"
            _ = fi15
        }
        var v18 float32 = (-2.25 * 3.14159)
        var v19 uint32 = genFn0(("" + "a"), (!(true && true)))
        _ = v18
        _ = v19
    }
    var v20 uint16 = (v9 | 255)
    v21 := []int{}
    if (556 != -1) {
        var v22 bool = (!(!false))
        v1[(("hello" + "a") + "foo bar")] = (("hello" + ("" + "hello")) + "\n\t")
        _ = v22
    }
    delete(v1, ("a" + ("a" + "a")))
    v23 := v1["hello"]
    for fi24 := 0; fi24 < 1; fi24++ {
        v25 := GenS2{F0: -337}
        delete(v1, v23)
        _ = fi24
        _ = v25
    }
    v9 = v9
    var v26 int64 = 42
    if false {
        for fi27 := 0; fi27 < 2; fi27++ {
            if true {
                _ = len(v21)
            }
            _ = fi27
        }
    } else {
        v23 = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v26
    _ = v3
    _ = v9
}
