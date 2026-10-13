package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 42
    var v3 string = ("foo bar" + "\n\t")
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 2; fi5++ {
        if false {
            var v6 int64 = 630
            _ = v6
        }
        _ = fi5
    }
    v7 := GenS1{F0: false, F1: 1}
    v8 := GenS1{F0: true, F1: 0}
    var v9 uint16 = 549
    for fi10 := 0; fi10 < 8; fi10++ {
        if (397 < 516) {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 bool = false
                v13 := []int{fi10, (fi11 >> ((0) & 63)), (fi10 << ((4294967295) & 63)), fi10}
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 int16 = (v2 % (((v2 | 1000)) | 1))
                v9 -= v9
                _ = fi14
                _ = v15
            }
        }
        switch "hello" {
        case "foo bar":
            for fi16 := 0; fi16 < 19; fi16++ {
                v2 = -844
                v17 := GenS1{F0: true, F1: 9223372036854775807}
                _ = fi16
                _ = v17
            }
        }
        _ = fi10
    }
    v3 = v3
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := []int{127, fi18, (fi18 >> ((258) & 63)), (-(fi18))}
        v20 := v1[(fi18 << ((380) & 63))]
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: 23}
    var v22 bool = false
    v23 := GenS2{F0: 1}
    v24 := GenS1{F0: (0 > 1), F1: -118}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
