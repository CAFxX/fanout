package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint32 = 0
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]bool)
    v7 := GenS1{F0: v1, F1: 757}
    v1 += 31
    if (4294967295 >= 227) {
        var v8 int8 = -87
        _ = v8
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 int64 = int64(v1)
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(fi11)
            switch "" {
            case "":
                switch (fi11 >> ((675) & 63)) {
                case -133:
                    var v12 int16 = -1
                    _ = v12
                case 721:
                    v1 -= 2147483647
                case -248:
                    _ = len(v3)
                }
            case "a":
                _ = len(v3)
            }
            _ = fi11
        }
        switch ("\n\t" + "") {
        case "":
            for fi13 := 0; fi13 < 8; fi13++ {
                v1 += (v1 / (((v1 & (v1 | (v1 + (v1 - v1))))) | 1))
                if (100 >= 100) {
                    fmt.Println(v1, v1)
                } else {
                    v1 = 452
                }
                switch (("" + ("" + "")) + "hello") {
                case "hello":
                    v1 -= 797
                case "\n\t":
                    for fi14 := 0; fi14 < 5; fi14++ {
                        if len(v2) > 0 {
                            v2[0] = 979
                        }
                        if false {
                            var v15 uint = 0
                            fmt.Println(fi14)
                            _ = v15
                        } else {
                            v1 = (v1 - v1)
                            var v16 float64 = 1e10
                            _ = v16
                        }
                        fmt.Println(fi13, fi13)
                        _ = fi14
                    }
                case "a":
                    fmt.Println(fi13)
                }
                _ = fi13
            }
        }
    }
    fmt.Println(v1)
    var v17 int = 10
    var v18 string = "foo bar"
    fmt.Println(v18, v17)
    fmt.Println(v17, v1, v1)
    var v19 int8 = 42
    if (true || true) {
        v18 = (("a" + v18) + v18)
        var v20 int32 = 784
        _ = v20
    } else {
        v19 = (v19 & 93)
    }
    _ = len(v2)
    if false {
        if false {
            var v21 uint32 = v1
            v22 := GenS1{F0: v21, F1: 2}
            _ = v21
            _ = v22
        }
        v18 = v18
        v23 := v6[("foo bar" + (("\n\t" + v18) + (("a" + v18) + v18)))]
        _ = v23
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            if false {
                var v25 int = fi24
                v26 := GenS2{F0: 888, F1: true}
                for fi27 := 0; fi27 < 13; fi27++ {
                    v1 = (^(v1))
                    _ = fi27
                }
                _ = v25
                _ = v26
            }
            v6[v18] = true
            _ = fi24
        }
        var v28 int16 = -393
        _ = len(v2)
        _ = v28
    }
    fmt.Println(v18, v18)
    var v29 int64 = -856
    v19 = 0
    var v30 int16 = 328
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
