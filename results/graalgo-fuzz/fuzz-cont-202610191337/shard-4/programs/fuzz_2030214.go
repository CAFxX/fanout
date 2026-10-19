package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[int]int)
    var v3 uint = 144
    v4 := make(map[uint32]bool)
    var v5 string = ((("a" + ("a" + "a")) + "a") + ("a" + "hello"))
    var v6 uint8 = 128
    fmt.Println(v1)
    if ((float32(1e10) - float32(1.5)) < ((0.0 - 3.14159) - (0.0 + 1.5))) {
        v7 := GenS1{F0: (-(v1)), F1: (100 == -30), F2: 42}
        var v8 float64 = (1.5 / 3.14159)
        _ = v7
        _ = v8
    } else {
        if true {
            v6 -= 255
            if false {
                switch -35 {
                case 0:
                    var v9 uint = (^(v3))
                    _ = v9
                case 736:
                    var v10 bool = false
                    _ = v10
                }
            }
            if (false && false) {
                for fi11 := 0; fi11 < 13; fi11++ {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        fmt.Println(v1, fi12)
                        var v13 uint64 = 551
                        v14 := GenS2{F0: 1}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    fmt.Println(v3, v3)
                    var v15 int16 = 1
                    _ = fi11
                    _ = v15
                }
                v16 := GenS2{F0: -760}
                var v17 bool = false
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 4; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    fmt.Println(fi19, v1)
                    switch fi18 {
                    case 0:
                        var v20 uint = v3
                        _ = v20
                    }
                    _ = fi19
                }
                var v21 uint8 = v6
                _ = fi18
                _ = v21
            }
            _ = len(v2)
            if false {
                switch "\n\t" {
                case "hello":
                    if (v6 <= 127) {
                        var v22 int32 = 0
                        for fi23 := 0; fi23 < 9; fi23++ {
                            var v24 string = (("" + "\n\t") + v5)
                            switch fi23 {
                            case -926:
                                if true {
                                    v25 := []int{-1, fi23}
                                    _ = v25
                                }
                            case -746:
                                v22 += (v22 << ((2) & 63))
                            case 597:
                                var v26 int64 = 9223372036854775807
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        v3 -= v3
                        _ = v22
                    }
                }
            } else {
                for fi27 := 0; fi27 < 8; fi27++ {
                    fmt.Println(v6, v3)
                    _ = fi27
                }
                fmt.Println(v1, v3, v3)
            }
        }
    }
    v28 := GenS2{F0: 1}
    var v29 float32 = 1e10
    var v30 uint64 = 4294967295
    var v31 string = ""
    v32 := []int{-1, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
