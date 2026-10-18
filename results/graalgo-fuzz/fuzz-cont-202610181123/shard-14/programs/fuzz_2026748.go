package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 int64; F1 int }

func main() {
    var v1 float64 = float64(0.0)
    var v2 uint = 0
    v3 := make(map[uint32]uint32)
    var v4 int8 = 112
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 float64 = v1
            v4 -= 0
            if false {
                v10 := GenS1{F0: 2, F1: 127}
                _ = v10
            } else {
                var v11 float32 = float32(v1)
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v12 uint64 = 1
    fmt.Println(v12, v4, v1)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int64 = 620
        var v15 bool = false
        fmt.Println(v15, v1)
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := []int{int(v12), 310}
            fmt.Println(v4, v2)
            _ = fi16
            _ = v17
        }
        fmt.Println(v2, v4)
    }
    v18 := GenS2{F0: -1, F1: -9223372036854775808, F2: 127}
    fmt.Println(v1, v12)
    v4 -= v4
    var v19 bool = true
    var v20 uint8 = 41
    for fi21 := 0; fi21 < 10; fi21++ {
        switch (fi21 - 0) {
        case -173:
            switch (("" + ("a" + ((("hello" + "a") + (("\n\t" + "") + "\n\t")) + ("hello" + "hello")))) + ("\n\t" + (("foo bar" + "") + ((("foo bar" + ("\n\t" + "")) + "\n\t") + "a")))) {
            case "\n\t":
                v22 := v6[2]
                _ = v22
            case "hello":
                var v23 int8 = (v4 | int8(v20))
                _ = v23
            }
        }
        var v24 string = ""
        v20 -= v20
        _ = fi21
        _ = v24
    }
    fmt.Println(v2, v1)
    for fi25 := 0; fi25 < 3; fi25++ {
        v26 := GenS3{F0: -9223372036854775808, F1: 0}
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if (2 >= 42) {
            v4 += (^(v4))
            fmt.Println(v2)
        }
        if (-329 == 1) {
            _ = len(v6)
            fmt.Println(v19, v4)
            var v28 bool = (255 == 904)
            _ = v28
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
