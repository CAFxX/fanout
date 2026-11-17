package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32, p1 int32) uint64 {
    var v6 uint32 = 2147483648
    var v7 string = (("\n\t" + "foo bar") + ("hello" + ""))
    for fi8 := 0; fi8 < 18; fi8++ {
        v7 = ""
        var v9 int = 55
        _ = fi8
        _ = v9
    }
    if true {
        p1 += (p1 % ((-184) | 1))
        v10 := GenS1{F0: -9223372036854775808}
        switch -1 {
        case 1:
            var v11 float32 = 1.5
            _ = v11
        }
        _ = v10
    }
    if (true && (true || (!(!(!(!true)))))) {
        var v12 uint32 = 613
        v13 := GenS1{F0: -9223372036854775808}
        v14 := []int{}
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 854
}

func main() {
    var v1 int16 = 0
    var v2 int8 = -1
    var v3 float64 = -2.25
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    switch (("a" + "a") + "hello") {
    case "a":
        _ = len(v5)
    case "foo bar":
        var v15 string = ("" + "hello")
        _ = v15
    case "\n\t":
        var v16 uint64 = 659
        _ = v16
    }
    var v17 bool = ((280 > 2147483647) && false)
    fmt.Println(v17)
    fmt.Println(v1, v1, v3)
    v18 := GenS1{F0: 1}
    if (!v17) {
        var v19 uint64 = 18446744073709551615
        if ((v3 / v3) > (((v3 / (v3 + (-2.25 - (float64(3.14159) * (v3 * v3))))) - 1.5) + ((v3 * 1e10) / ((v3 * -2.25) - float64(0.1))))) {
            var v20 float64 = v3
            var v21 string = "foo bar"
            _ = v20
            _ = v21
        }
        var v22 int64 = 505
        _ = v19
        _ = v22
    }
    v17 = (32767 < 65535)
    switch -2147483648 {
    case 1:
        v2 = (v2 - v2)
    case 127:
        var v23 int8 = (^(v2))
        _ = v23
    }
    switch "a" {
    case "hello":
        v2 = (v2 ^ v2)
    case "foo bar":
        v17 = v17
    }
    _ = len(v4)
    if v17 {
        var v24 int64 = 680
        _ = v24
    } else {
        v25 := []int{}
        v17 = v17
        var v26 int32 = int32(v2)
        _ = v25
        _ = v26
    }
    var v27 uint16 = 986
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
