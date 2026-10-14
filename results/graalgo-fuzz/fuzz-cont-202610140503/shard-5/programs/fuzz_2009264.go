package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16) uint {
    fmt.Println(p0)
    var v7 bool = (false && (127 <= 288))
    _ = p0
    _ = v7
    return 605
}
func genFn1() int {
    var v8 int = 243
    v9 := []int{v8}
    switch (("\n\t" + ("" + (("hello" + "a") + ""))) + ((("\n\t" + "a") + ("" + (("a" + "\n\t") + "hello"))) + ("hello" + "foo bar"))) {
    case "foo bar":
        fmt.Println(v8)
    }
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 bool = (133 >= 2)
    var v2 uint8 = 175
    var v3 int32 = 42
    v4 := []int{}
    var v5 uint32 = 4294967295
    var v6 int64 = -9223372036854775808
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 bool = false
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(v2, fi12, v11)
            v1 = true
            _ = fi12
        }
        var v13 int64 = v6
        _ = fi10
        _ = v11
        _ = v13
    }
    if false {
        v6 = -1
        var v14 float32 = 3.14159
        _ = v14
    }
    var v15 int32 = v3
    if true {
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint16 = 3
                fmt.Println(v3, fi16)
                _ = fi17
                _ = v18
            }
            if v1 {
                var v19 uint8 = uint8(v5)
                _ = v19
            }
            v20 := GenS1{F0: 100}
            _ = fi16
            _ = v20
        }
        var v21 float64 = 0.1
        _ = v21
    } else {
        var v22 float64 = 3.14159
        _ = v22
    }
    var v23 int32 = v3
    var v24 uint8 = v2
    v15 = (v3 | (v3 | -825))
    var v25 int64 = (v6 * v6)
    var v26 uint = 732
    v24 = (v2 % ((v2) | 1))
    v24 -= v2
    v24 += 2
    v27 := GenS1{F0: -123}
    v2 = v24
    if (!true) {
        var v28 int64 = int64(v5)
        v29 := GenS1{F0: 0}
        _ = v28
        _ = v29
    } else {
        if len(v4) > 0 {
            v4[0] = genFn1()
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
