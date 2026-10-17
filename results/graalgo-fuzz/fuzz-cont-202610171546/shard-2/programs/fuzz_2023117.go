package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 int32; F2 uint32 }

func main() {
    var v1 uint64 = 1
    var v2 uint = 42
    v3 := make(map[int]int)
    var v4 float64 = 3.14159
    var v5 uint = v2
    v6 := []int{-20, -469}
    v2 -= (v5 - (v5 * uint(v1)))
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int8 = 1
        var v9 float64 = v4
        var v10 uint16 = 395
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if (!true) {
        switch ("\n\t" + "") {
        case "\n\t":
            for fi11 := 0; fi11 < 5; fi11++ {
                if (true || true) {
                    for fi12 := 0; fi12 < 10; fi12++ {
                        fmt.Println(fi12, fi12, v2)
                        _ = fi12
                    }
                    switch fi11 {
                    case 32767:
                        if true {
                            fmt.Println(v2)
                            v13 := GenS1{F0: 146}
                            _ = v13
                        }
                    case 1000:
                        var v14 int8 = -1
                        _ = v14
                    }
                } else {
                    var v15 string = ("hello" + "a")
                    _ = v15
                }
                if true {
                    v5 = 42
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v5 += (v2 / ((v5) | 1))
                        _ = fi16
                    }
                }
                _ = fi11
            }
        }
        v2 = 255
    }
    v17 := GenS1{F0: 42}
    v18 := GenS2{F0: 255}
    v5 = v5
    fmt.Println(v2)
    v5 += v2
    var v19 uint32 = 958
    fmt.Println(v5)
    var v20 uint8 = 128
    fmt.Println(v1, v1)
    v21 := v3[1]
    for fi22 := 0; fi22 < 6; fi22++ {
        if ((!(0 == 127)) || false) {
            if true {
                for fi23 := 0; fi23 < 7; fi23++ {
                    v5 = (v2 * v5)
                    if false {
                        if true {
                            v20 -= v20
                        }
                    }
                    var v24 bool = (460 > 578)
                    _ = fi23
                    _ = v24
                }
                switch fi22 {
                case 4294967295:
                    v2 = 255
                case -1:
                    v5 = 346
                }
            } else {
                v25 := GenS2{F0: v20}
                fmt.Println(fi22)
                var v26 bool = (v21 == (v21 * fi22))
                _ = v25
                _ = v26
            }
            var v27 string = "hello"
            _ = v27
        } else {
            for fi28 := 0; fi28 < 3; fi28++ {
                v29 := GenS3{F0: v2, F1: -876, F2: v19}
                _ = fi28
                _ = v29
            }
        }
        var v30 uint = (v5 << (((^(v2))) & 63))
        _ = fi22
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
