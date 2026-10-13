package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32, p1 uint8) uint16 {
    switch "a" {
    case "a":
        fmt.Println(p0)
    default:
        for fi4 := 0; fi4 < 11; fi4++ {
            v5 := []int{0}
            _ = fi4
            _ = v5
        }
    }
    if true {
        p1 -= p1
    }
    var v6 string = "\n\t"
    _ = p0
    _ = p1
    _ = v6
    return 416
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 65535
    v7 := GenS1{F0: 2147483647, F1: 768, F2: "\n\t"}
    v3 = (v3 << ((4294967295) & 63))
    if (true && (!(false || (65535 == 255)))) {
        var v8 uint16 = 2
        v9 := GenS1{F0: 42, F1: 1, F2: ""}
        for fi10 := 0; fi10 < 3; fi10++ {
            switch ("\n\t" + ("\n\t" + "")) {
            case "foo bar":
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 int8 = 1
                    var v13 string = "foo bar"
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                v3 -= (v3 + (v3 / (((^(v3))) | 1)))
            }
            var v14 int16 = 1000
            _ = fi10
            _ = v14
        }
        _ = v8
        _ = v9
    }
    v3 = v3
    var v15 string = "\n\t"
    for fi16 := 0; fi16 < 19; fi16++ {
        v17 := GenS2{F0: -1, F1: 0}
        if false {
            var v18 uint64 = 4294967295
            if false {
                switch (fi16 & fi16) {
                case 42:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v15 = v15
                        switch -584 {
                        case 4294967295:
                            var v20 uint32 = (v3 % (((v3 - (^(v3)))) | 1))
                            _ = v20
                        default:
                            v15 = "foo bar"
                        }
                        var v21 uint = 821
                        _ = fi19
                        _ = v21
                    }
                }
                var v22 float64 = ((1.5 - 0.0) + (((1.5 / 1.5) - 0.1) * ((-2.25 - ((1e10 * 1e10) - -2.25)) / 1e10)))
                var v23 uint8 = 42
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: false}
            _ = v18
            _ = v24
        } else {
            if true {
                if ((fi16 + (fi16 + fi16)) <= 32768) {
                    v25 := []int{fi16, fi16}
                    var v26 int64 = 0
                    v26 = (v26 ^ v26)
                    _ = v25
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 18; fi27++ {
                    switch (v15 + v15) {
                    case "\n\t":
                        v28 := GenS1{F0: -659, F1: 1, F2: "\n\t"}
                        _ = v28
                    }
                    _ = len(v1)
                    _ = fi27
                }
                var v29 uint8 = 211
                v30 := GenS2{F0: 0, F1: 1}
                _ = v29
                _ = v30
            }
            if true {
                var v31 int8 = 15
                _ = v31
            } else {
                var v32 float64 = 0.1
                v32 += v32
                _ = v32
            }
            var v33 int64 = 42
            _ = v33
        }
        _ = fi16
        _ = v17
    }
    var v34 uint16 = genFn0(-1, 0)
    fmt.Println(v3, v3, v34)
    var v35 uint32 = v3
    v36 := []int{-9223372036854775808, -1, -818, 9223372036854775807}
    var v37 uint32 = (v3 * v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v7
}
