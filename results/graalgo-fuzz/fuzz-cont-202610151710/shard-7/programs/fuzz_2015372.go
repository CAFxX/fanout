package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []int{}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 string = "hello"
    var v5 uint8 = 255
    var v6 float64 = -2.25
    for fi7 := 0; fi7 < 12; fi7++ {
        if ("hello" != "a") {
            v4 = "a"
        } else {
            var v8 int64 = 42
            v4 = v4
            _ = v8
        }
        v2[(v4 + (v4 + v4))] = fi7
        var v9 int64 = -109
        _ = fi7
        _ = v9
    }
    v4 = v4
    var v10 uint16 = 826
    if (-1 == 42) {
        if true {
            v11 := GenS1{F0: -537}
            v12 := GenS1{F0: 772}
            for fi13 := 0; fi13 < 19; fi13++ {
                for fi14 := 0; fi14 < 16; fi14++ {
                    delete(v3, 760)
                    v4 = "\n\t"
                    var v15 string = "foo bar"
                    _ = fi14
                    _ = v15
                }
                var v16 int8 = 127
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int16 = 32767
                    v19 := GenS1{F0: 9223372036854775807}
                    v20 := GenS1{F0: 9223372036854775807}
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                _ = fi13
                _ = v16
            }
            _ = v11
            _ = v12
        } else {
            var v21 int64 = 1
            _ = v21
        }
        var v22 bool = (false && false)
        _ = v22
    } else {
        var v23 bool = true
        _ = v23
    }
    var v24 bool = false
    if v24 {
        fmt.Println(v24)
    } else {
        v25 := GenS1{F0: 801}
        _ = v25
    }
    v1 = append(v1, -874)
    if false {
        if true {
            v6 = v6
        }
    }
    var v26 int32 = -906
    if (!((v26 / ((543) | 1)) == -2147483648)) {
        var v27 int8 = 127
        _ = v27
    }
    var v28 int = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
