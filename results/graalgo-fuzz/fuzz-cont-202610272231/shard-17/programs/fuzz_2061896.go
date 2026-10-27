package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0() uint32 {
    if false {
        var v4 bool = false
        for fi5 := 0; fi5 < 9; fi5++ {
            switch "foo bar" {
            case "":
                v6 := []int{fi5, fi5, fi5}
                _ = v6
            case "hello":
                v7 := []int{fi5, fi5, 4294967295}
                _ = v7
            case "foo bar":
                fmt.Println(fi5)
            }
            v8 := []int{fi5}
            _ = fi5
            _ = v8
        }
        _ = v4
    } else {
        var v9 int8 = 100
        v9 = v9
        var v10 int32 = -2147483648
        _ = v10
        _ = v9
    }
    var v11 uint64 = 18446744073709551615
    _ = v11
    return 981
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42}
    var v3 int32 = -2147483648
    var v12 int8 = 127
    if (0 < 864) {
        var v13 uint32 = 711
        v2 = append(v2, (v13 >> ((321) & 63)))
        switch "a" {
        case "foo bar":
            var v14 int16 = 1000
            _ = v14
        case "\n\t":
            v15 := v1[128]
            _ = v15
        case "a":
            if len(v2) > 0 {
                v2[0] = (^(v13))
            }
        }
        _ = v13
    } else {
        for fi16 := 0; fi16 < 19; fi16++ {
            v17 := GenS1{F0: 42, F1: 4294967295, F2: 0}
            var v18 uint64 = 2
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v3 += (v3 - 132)
    var v19 bool = false
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v1[9223372036854775807] = 4294967295
    } else {
        var v20 float32 = -2.25
        fmt.Println(v20, v20)
        _ = v20
    }
    if v19 {
        v21 := GenS1{F0: 255, F1: 639, F2: -9223372036854775808}
        _ = v21
    }
    if false {
        var v22 int = 4294967295
        v23 := GenS1{F0: 42, F1: 65535, F2: -144}
        var v24 uint8 = 255
        _ = v22
        _ = v23
        _ = v24
    } else {
        var v25 int = 253
        if false {
            fmt.Println(v25)
            if (1.5 > -2.25) {
                fmt.Println(v3)
                var v26 int8 = 127
                _ = v26
            }
        }
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
}
