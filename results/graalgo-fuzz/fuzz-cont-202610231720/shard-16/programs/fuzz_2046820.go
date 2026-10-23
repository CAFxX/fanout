package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 uint8) int64 {
    v6 := GenS1{F0: 127}
    var v7 int32 = 42
    var v8 int64 = 1
    var v9 float64 = (-2.25 / (0.1 * 0.1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 9223372036854775807
}
func genFn1() uint64 {
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint = 862
            v13 := []int{fi10}
            v12 -= 4294967295
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    var v14 int64 = 1
    _ = v14
    return 2
}

func main() {
    var v1 float32 = (1.5 * ((1e10 * (1.5 * 3.14159)) / 0.1))
    var v2 uint32 = 71
    var v3 uint32 = v2
    var v4 string = ("a" + "a")
    var v5 int16 = -32768
    var v15 uint8 = 221
    for fi16 := 0; fi16 < 15; fi16++ {
        v2 -= (v2 | (v2 ^ 4294967295))
        var v17 uint = 244
        _ = fi16
        _ = v17
    }
    if false {
        v1 = ((((v1 * v1) / 1.5) - float32(1e10)) - v1)
        v18 := GenS1{F0: 2147483647}
        _ = v18
    } else {
        fmt.Println(v3)
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 int = (fi19 | fi19)
            if false {
                switch v4 {
                case "hello":
                    switch (v4 + "a") {
                    case "hello":
                        fmt.Println(fi19, v5)
                    }
                }
            } else {
                var v21 int16 = v5
                fmt.Println(v15, v4)
                fmt.Println(v1)
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        var v22 int16 = v5
        _ = v22
    }
    var v23 int = 127
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 uint16 = 65535
        v5 = v5
        fmt.Println(v25, v4, v3)
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        switch (v23 - fi26) {
        case 838:
            var v27 string = v4
            _ = v27
        default:
            var v28 uint32 = 4294967295
            _ = v28
        }
        _ = fi26
    }
    v29 := []int{v23}
    if true {
        v29 = append(v29, v23)
        var v30 int16 = (^(v5))
        _ = v30
    }
    var v31 uint8 = v15
    var v32 uint8 = v15
    switch "foo bar" {
    case "hello":
        switch (-(v23)) {
        case 184:
            v23 = v23
        case -1:
            v31 = (^(v31))
        case -408:
            if (1 != 42) {
                fmt.Println(v23)
            }
        }
    case "a":
        var v33 uint64 = 247
        _ = v33
    case "\n\t":
        var v34 int = v23
        _ = v34
    default:
        if true {
            var v35 int8 = 1
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
