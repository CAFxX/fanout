package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]bool)
    var v3 string = (v1 + ("\n\t" + ((v1 + "\n\t") + "\n\t")))
    var v4 int = -220
    for fi5 := 0; fi5 < 3; fi5++ {
        switch v1 {
        case "foo bar":
            if true {
                v6 := GenS1{F0: -1}
                for fi7 := 0; fi7 < 5; fi7++ {
                    var v8 uint64 = 18446744073709551615
                    _ = fi7
                    _ = v8
                }
                v3 = v1
                _ = v6
            } else {
                if (865 <= -1) {
                    var v9 string = v1
                    fmt.Println(v9)
                    switch v3 {
                    case "hello":
                        delete(v2, 163)
                    }
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        var v11 uint8 = 0
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS2{F0: "\n\t", F1: v3}
                    _ = v12
                }
            }
        case "\n\t":
            if (97 != 4294967295) {
                v3 = "\n\t"
                v4 = 335
            } else {
                fmt.Println(v4)
                v13 := GenS2{F0: v1, F1: "hello"}
                v1 = "\n\t"
                _ = v13
            }
        }
        var v14 int32 = -2147483648
        _ = fi5
        _ = v14
    }
    fmt.Println(v1, v4)
    var v15 float64 = 1.5
    var v16 uint32 = 2147483648
    v16 -= v16
    var v17 string = (v3 + v3)
    var v18 uint32 = v16
    var v19 int8 = -128
    for fi20 := 0; fi20 < 17; fi20++ {
        var v21 float64 = (0.0 + float64(1.5))
        _ = fi20
        _ = v21
    }
    switch (v4 & (v4 & 128)) {
    case 32768:
        var v22 float32 = -2.25
        _ = v22
    default:
        var v23 string = (v1 + "\n\t")
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 int8 = v19
            var v27 int32 = 2147483647
            _ = fi25
            _ = v26
            _ = v27
        }
        if false {
            v16 += (v18 << ((451) & 63))
            if true {
                v28 := []int{0, int(v19), 4294967295, v4}
                if false {
                    v16 = (v18 >> ((2) & 63))
                    for fi29 := 0; fi29 < 12; fi29++ {
                        fmt.Println(v18, v18, v17)
                        fmt.Println(v4, v18)
                        fmt.Println(v16)
                        var v30 int32 = 1
                        _ = fi29
                        _ = v30
                    }
                }
                for fi31 := 0; fi31 < 18; fi31++ {
                    fmt.Println(v16, v18)
                    v4 -= (fi31 & (fi24 >> ((1) & 63)))
                    _ = fi31
                }
                _ = v28
            }
        } else {
            if false {
                var v32 int = (^(v4))
                v18 = v16
                _ = v32
            }
            v18 = v16
        }
        var v33 int32 = 449
        _ = fi24
        _ = v33
    }
    var v34 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
