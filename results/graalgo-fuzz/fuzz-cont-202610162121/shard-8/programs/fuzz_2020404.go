package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
func genFn0() int64 {
    v6 := GenS1{F0: 2}
    var v7 string = ""
    _ = v6
    _ = v7
    return 696
}

func main() {
    var v1 uint8 = 6
    v2 := make(map[int]uint32)
    var v3 uint8 = v1
    v4 := make(map[string]string)
    v5 := []string{}
    for fi8 := 0; fi8 < 9; fi8++ {
        _ = len(v5)
        _ = fi8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS2{F0: "a"}
        _ = fi9
        _ = v10
    }
    v1 += 255
    var v11 uint8 = v1
    var v12 uint32 = 65535
    for fi13 := 0; fi13 < 9; fi13++ {
        fmt.Println(v3, v12, fi13)
        _ = fi13
    }
    var v14 uint32 = 1
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 5; fi16++ {
        if false {
            var v17 int8 = 0
            fmt.Println(v3)
            _ = v17
        } else {
            for fi18 := 0; fi18 < 5; fi18++ {
                if (true || (!false)) {
                    v3 = (v3 + 2)
                    v12 = v12
                }
                var v19 int16 = 42
                v11 += (v1 >> ((42) & 63))
                if false {
                    v2[349] = (v12 / ((4294967295) | 1))
                    v12 = (v12 - (v14 % ((4294967295) | 1)))
                    var v20 uint64 = 2
                    _ = v20
                } else {
                    fmt.Println(v1)
                }
                _ = fi18
                _ = v19
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                v14 = (v12 - (v12 % (((v14 + v12)) | 1)))
                var v22 int64 = genFn0()
                var v23 uint8 = v1
                v24 := GenS3{F0: 50}
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            v1 += 44
        }
        var v25 uint16 = 42
        var v26 uint16 = 2
        _ = fi16
        _ = v25
        _ = v26
    }
    var v27 int = -2147483648
    for fi28 := 0; fi28 < 10; fi28++ {
        v29 := GenS3{F0: (v3 % ((v11) | 1))}
        _ = fi28
        _ = v29
    }
    fmt.Println(v1, v1, v1)
    if false {
        var v30 uint = 18446744073709551615
        v31 := GenS4{F0: 1}
        _ = v30
        _ = v31
    }
    var v32 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
