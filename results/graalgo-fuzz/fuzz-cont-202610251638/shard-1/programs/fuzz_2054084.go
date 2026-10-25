package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint64, p1 int64, p2 int16) uint16 {
    var v5 uint8 = 2
    var v6 uint32 = 2
    if true {
        v7 := GenS1{F0: v5}
        fmt.Println(p2, p1)
        _ = v7
    } else {
        v8 := GenS1{F0: (v5 ^ v5)}
        _ = v8
    }
    v9 := []int{4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := []int{1, 128, 42, -9223372036854775808}
    var v3 uint64 = 821
    var v4 uint = 42
    var v10 string = (("\n\t" + "a") + "")
    v10 = "foo bar"
    for fi11 := 0; fi11 < 1; fi11++ {
        v1[(-(fi11))] = v10
        _ = fi11
    }
    v4 = uint(v3)
    var v12 string = ""
    if (255 < 65535) {
        v3 += (^(v3))
    } else {
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    delete(v1, 9223372036854775807)
    switch ("\n\t" + (v10 + "foo bar")) {
    case "hello":
        v14 := v1[32768]
        _ = v14
    case "a":
        var v15 uint32 = 65535
        _ = v15
    }
    if ((uint8(v3) != 116) || ((float64((1e10 + (3.14159 * (1.5 / float32((1.5 + ((((0.1 - float64((float32(0.1) * 1e10))) + 0.0) - 1.5) * 0.0))))))) + -2.25) > float64(0.0))) {
        fmt.Println(v4, v3, v4)
    }
    var v16 uint16 = genFn0(42, 42, 284)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
