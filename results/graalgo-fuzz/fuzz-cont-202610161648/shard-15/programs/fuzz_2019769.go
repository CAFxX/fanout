package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
func genFn0(p0 uint16, p1 uint8) int64 {
    p0 = (p0 % ((p0) | 1))
    var v4 int16 = 1
    _ = p0
    _ = p1
    _ = v4
    return 687
}
func genFn1(p0 uint64, p1 int32, p2 uint64) uint16 {
    p2 = (^(p0))
    switch ("a" + "hello") {
    case "a":
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := []int{(fi5 << ((2) & 63)), -1000, (fi5 << ((18446744073709551615) & 63)), 673}
            if (true || ((1.5 != 0.0) && (90 >= uint16(p0)))) {
                fmt.Println(p1, p0, p0)
                fmt.Println(p1, p1, p2)
                for fi7 := 0; fi7 < 7; fi7++ {
                    var v8 int32 = p1
                    _ = fi7
                    _ = v8
                }
            } else {
                v9 := GenS1{F0: -122}
                v10 := GenS1{F0: 127}
                var v11 bool = (true || false)
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := GenS1{F0: 49}
            _ = fi5
            _ = v12
            _ = v6
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int16 = 42
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]int)
    fmt.Println(v2)
    v1 = v1
    for fi13 := 0; fi13 < 11; fi13++ {
        if (true && (!true)) {
            var v14 int32 = -126
            _ = v14
        }
        _ = fi13
    }
    v15 := GenS2{F0: -56, F1: 0, F2: 2147483647}
    switch "\n\t" {
    case "\n\t":
        v1 += (v1 & (v1 / ((v1) | 1)))
    case "":
        for fi16 := 0; fi16 < 12; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 int16 = (v1 & 42)
                _ = fi17
                _ = v18
            }
            v19 := v3[(v2 / ((705) | 1))]
            v19 = v19
            _ = fi16
            _ = v19
        }
    }
    if true {
        v2 = 1
    } else {
        switch (("hello" + "") + "hello") {
        case "a":
            v2 = v2
        }
    }
    var v20 bool = (100 == -128)
    if ((0 > 0) || v20) {
        v2 = (v2 >> ((748) & 63))
        v21 := v3[(^(v2))]
        var v22 uint8 = 42
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
