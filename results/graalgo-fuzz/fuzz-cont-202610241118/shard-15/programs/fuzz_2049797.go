package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint = 356
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    var v4 float64 = float64((1.5 / (1e10 + float32((1.5 - -2.25)))))
    var v5 uint16 = 0
    var v6 float64 = ((-2.25 - (3.14159 + -2.25)) / 0.1)
    v7 := GenS1{F0: 42, F1: -223}
    for fi8 := 0; fi8 < 3; fi8++ {
        switch ("" + "hello") {
        case "hello":
            v4 -= v6
        case "":
            var v9 uint8 = 11
            _ = v9
        default:
            for fi10 := 0; fi10 < 19; fi10++ {
                v5 = v5
                v11 := GenS1{F0: 0, F1: -1}
                var v12 string = ""
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        _ = fi8
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 42, F1: -455}
        var v15 int16 = -368
        if true {
            switch (fi13 + (-(fi13))) {
            case 0:
                if false {
                    delete(v3, 2147483647)
                    switch "hello" {
                    case "\n\t":
                        v15 -= (v15 - 1)
                    case "foo bar":
                        fmt.Println(v4)
                    }
                    switch "hello" {
                    case "\n\t":
                        v16 := GenS1{F0: int8(v1), F1: 1}
                        _ = v16
                    case "a":
                        delete(v3, 468)
                    }
                } else {
                    var v17 uint = 0
                    fmt.Println(v6, v1, v1)
                    _ = v17
                }
            case 42:
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (!(!true)) {
                        v5 = v5
                        v19 := GenS1{F0: 97, F1: -1}
                        v20 := GenS1{F0: 42, F1: -2147483648}
                        _ = v19
                        _ = v20
                    } else {
                        v15 = v15
                    }
                    var v21 uint = (^(v1))
                    _ = fi18
                    _ = v21
                }
            case -319:
                var v22 int32 = 42
                _ = v22
            }
            var v23 uint16 = v5
            var v24 uint = v1
            _ = v23
            _ = v24
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if (128 <= 0) {
        switch -1 {
        case 356:
            fmt.Println(v6, v6, v6)
        case -2147483648:
            var v25 int32 = 2147483647
            _ = v25
        }
        v5 = (v5 & v5)
    }
    if (!((18446744073709551615 != 42) && (2 != uint8(v1)))) {
        v26 := []int{127, 295}
        fmt.Println(v6, v6, v5)
        v1 -= (v1 | (v1 / ((v1) | 1)))
        _ = v26
    }
    var v27 uint64 = 42
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v5, fi28)
        _ = fi28
    }
    v29 := GenS1{F0: 42, F1: 2147483647}
    v1 += (^(v1))
    v5 = v5
    fmt.Println(v27, v6)
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 uint8 = 42
        if true {
            var v32 uint8 = v31
            v1 += (^(v1))
            _ = v32
        } else {
            var v33 uint32 = 423
            v34 := []int{(fi30 / ((fi30) | 1))}
            _ = v33
            _ = v34
        }
        _ = fi30
        _ = v31
    }
    var v35 uint = (v1 >> ((v1) & 63))
    for fi36 := 0; fi36 < 2; fi36++ {
        if (v27 < (v27 + (v27 % (((^(v27))) | 1)))) {
            var v37 uint8 = uint8(v35)
            v38 := []int{}
            var v39 uint64 = (v27 >> ((506) & 63))
            _ = v37
            _ = v38
            _ = v39
        }
        v6 = (0.0 - (0.1 - v6))
        v35 += (v1 / ((v1) | 1))
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
