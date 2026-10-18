package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    var v3 float64 = float64(1e10)
    v4 := make(map[int]int)
    var v5 float32 = (((3.14159 * ((1e10 / (0.0 - -2.25)) / 1e10)) * 3.14159) - ((1e10 / float32((v3 / ((v3 * 0.1) - 0.0)))) - float32((v3 - ((v3 * 1e10) - float64(-2.25))))))
    var v6 bool = true
    fmt.Println(v5, v6)
    switch ("" + "\n\t") {
    case "\n\t":
        var v7 int64 = 42
        _ = v7
    default:
        v3 -= v3
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v5, v3, v3)
        if (42 <= 9223372036854775807) {
            var v9 uint64 = 42
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int32 = -2147483648
                v5 -= 1.5
                _ = fi10
                _ = v11
            }
            _ = v9
        } else {
            switch (fi8 << ((4294967295) & 63)) {
            case 774:
                v12 := []int{fi8}
                _ = v12
            }
            var v13 int8 = 49
            _ = v13
        }
        _ = fi8
    }
    v14 := GenS1{F0: 65535}
    var v15 uint32 = 1
    v15 -= v15
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 uint = 65535
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 float64 = (v3 * v3)
            v5 = v5
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        switch (fi20 - (fi20 % (((fi20 * fi20)) | 1))) {
        case 181:
            v15 = (^(v15))
        }
        _ = fi20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v3 = 0.0
        v5 = -2.25
        v15 -= (^(v15))
        _ = fi21
    }
    var v22 float32 = (v5 * v5)
    v23 := GenS1{F0: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
