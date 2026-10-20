package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 int = 626
    v5 := make(map[string]bool)
    v6 := make(map[string]string)
    var v7 int16 = (v1 - 42)
    for fi8 := 0; fi8 < 19; fi8++ {
        if v3 {
            var v9 uint32 = 494
            _ = v9
        } else {
            v10 := []int{v4}
            var v11 bool = true
            _ = v10
            _ = v11
        }
        if false {
            if v3 {
                fmt.Println(fi8)
                fmt.Println(v4)
            }
        } else {
            v12 := []int{}
            _ = len(v5)
            v13 := GenS1{F0: 399}
            _ = v12
            _ = v13
        }
        var v14 uint16 = 42
        _ = fi8
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        switch (fi15 / ((v4) | 1)) {
        case 4294967295:
            fmt.Println(fi15, v3)
        case 42:
            var v16 int8 = 75
            _ = v16
        }
        v4 = (-(v4))
        var v17 uint16 = 42
        _ = fi15
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 bool = false
        var v20 int32 = -914
        var v21 uint64 = 18446744073709551615
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 int32 = 1
    v3 = false
    v23 := GenS1{F0: 567}
    if (!false) {
        if (!true) {
            v3 = (!v3)
            var v24 int8 = 127
            if v3 {
                var v25 int32 = v22
                var v26 bool = (!(false || (!(((192 == v4) && ((v24 % ((42) | 1)) == 5)) && v3))))
                v7 += (v1 / ((v7) | 1))
                _ = v25
                _ = v26
            } else {
                v27 := GenS1{F0: 2}
                _ = v27
            }
            _ = v24
        }
        v22 = (v22 | 585)
        fmt.Println(v4)
    }
    var v28 int64 = 1
    var v29 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
