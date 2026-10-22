package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := []uint32{v1}
    var v5 string = "foo bar"
    v6 := []int{67}
    var v7 bool = (954 > 32767)
    var v8 int8 = 43
    var v9 string = v5
    _ = len(v4)
    var v10 int8 = v8
    fmt.Println(v5, v1)
    if true {
        v9 = (v9 + v5)
    }
    delete(v3, 442)
    if v7 {
        if len(v4) > 0 {
            v4[0] = v1
        }
        var v11 int16 = 32767
        _ = v11
    }
    v1 -= (v1 << ((2) & 63))
    var v12 uint32 = 1
    if v7 {
        var v13 int8 = v10
        if true {
            if (6 > 42) {
                if (v8 == (v13 | -20)) {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v2[(-(fi14))] = v12
                        var v15 float64 = 1e10
                        _ = fi14
                        _ = v15
                    }
                    if false {
                        v12 = v12
                        fmt.Println(v12)
                        v5 = v9
                    }
                }
            } else {
                if (-1 < 0) {
                    if (2 > 1) {
                        var v16 uint = 65535
                        _ = v16
                    }
                }
                var v17 int8 = v10
                _ = v17
            }
            var v18 uint16 = 663
            v13 += v8
            _ = v18
        }
        switch (v9 + v5) {
        case "":
            v19 := GenS1{F0: "hello"}
            _ = v19
        }
        _ = v13
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 float32 = 1.5
        if (float64(0.0) < float64(1.5)) {
            v22 := GenS1{F0: "\n\t"}
            delete(v2, int(v8))
            _ = v22
        } else {
            if v7 {
                v23 := GenS1{F0: "a"}
                if (0 > 42) {
                    fmt.Println(v8, v12, v21)
                }
                _ = v23
            }
            v24 := GenS1{F0: "hello"}
            var v25 int8 = 127
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    v5 = v9
    v26 := GenS1{F0: (v5 + "a")}
    var v27 uint32 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
