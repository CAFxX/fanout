package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int32 = -1
    var v2 float32 = (0.1 + 1e10)
    v3 := make(map[uint32]bool)
    switch "a" {
    case "foo bar":
        for fi4 := 0; fi4 < 6; fi4++ {
            var v5 int32 = 0
            v6 := []int{fi4}
            if false {
                v3[2147483648] = (!true)
                switch "hello" {
                case "hello":
                    v5 = 552
                case "foo bar":
                    if false {
                        v2 += 3.14159
                        if (128 == 1) {
                            v2 = ((float32(0.0) / (v2 * v2)) - v2)
                        } else {
                            for fi7 := 0; fi7 < 11; fi7++ {
                                v8 := GenS1{F0: v5}
                                if true {
                                    var v9 int16 = -504
                                    var v10 uint8 = 42
                                    _ = v10
                                    _ = v9
                                }
                                _ = fi7
                                _ = v8
                            }
                        }
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        var v11 uint64 = 0
                        _ = v11
                    }
                }
                var v12 uint64 = 59
                _ = v12
            }
            _ = fi4
            _ = v5
            _ = v6
        }
    case "":
        var v13 int32 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: (^(v1))}
        var v16 float32 = v2
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 += (1e10 + (v2 - v2))
    _ = len(v3)
    fmt.Println(v2, v2)
    if true {
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS1{F0: (v1 << ((255) & 63))}
            fmt.Println(v2, v2, fi17)
            _ = fi17
            _ = v18
        }
        if true {
            if false {
                var v19 int32 = 1
                if false {
                    var v20 int32 = 0
                    v3[459] = false
                    _ = v20
                } else {
                    v19 = v19
                }
                _ = v19
            } else {
                v2 -= float32(1e10)
                fmt.Println(v1)
            }
        }
    } else {
        v1 += (v1 >> ((255) & 63))
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if false {
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 bool = false
                var v24 uint8 = 2
                fmt.Println(fi22)
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            switch (-(fi21)) {
            case 2147483647:
                var v25 uint32 = 0
                _ = v25
            case 127:
                var v26 int32 = (^(v1))
                _ = v26
            default:
                var v27 uint8 = 2
                _ = v27
            }
            v1 = (v1 ^ v1)
        }
        _ = fi21
    }
    v28 := GenS1{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
}
