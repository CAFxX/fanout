package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS5 struct { F0 int16 }

func main() {
    var v1 int8 = 100
    var v2 uint64 = 18446744073709551615
    var v3 string = ((("\n\t" + "a") + ("\n\t" + "")) + "hello")
    v4 := make(map[uint32]int)
    var v5 int16 = -248
    v6 := make(map[string]int)
    v7 := v4[65535]
    fmt.Println(v7)
    var v8 uint16 = 42
    v9 := GenS1{F0: 0}
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            switch "foo bar" {
            case "foo bar":
                var v12 int16 = (v5 >> ((42) & 63))
                _ = v12
            case "a":
                switch fi10 {
                case 71:
                    fmt.Println(fi11, v3)
                }
            }
            v13 := GenS2{F0: -1, F1: 0}
            var v14 uint = 677
            _ = fi11
            _ = v13
            _ = v14
        }
        _ = fi10
    }
    v15 := v6[(v3 + "")]
    var v16 int16 = v5
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 int = (fi18 & (fi17 & -2147483648))
            switch v3 {
            case "\n\t":
                v1 = (v1 << ((2) & 63))
            default:
                v20 := GenS3{F0: 128}
                _ = v20
            }
            var v21 int8 = (^(v1))
            _ = fi18
            _ = v19
            _ = v21
        }
        _ = fi17
    }
    v22 := GenS4{F0: 434, F1: -1, F2: 0}
    if false {
        var v23 int16 = -764
        if (v1 <= -128) {
            v24 := v6["hello"]
            v7 += v24
            _ = v24
        }
        _ = v23
    } else {
        if (false && true) {
            if (255 == 580) {
                if false {
                    var v25 uint = 811
                    switch ("a" + (v3 + "\n\t")) {
                    case "":
                        if (91 >= 4294967295) {
                            for fi26 := 0; fi26 < 20; fi26++ {
                                if true {
                                    var v27 uint32 = 2147483648
                                    v15 = int(v8)
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 11; fi28++ {
                                    v29 := []int{}
                                    _ = fi28
                                    _ = v29
                                }
                                _ = fi26
                            }
                            var v30 uint16 = 1
                            for fi31 := 0; fi31 < 18; fi31++ {
                                var v32 uint = v25
                                _ = fi31
                                _ = v32
                            }
                            _ = v30
                        }
                    case "a":
                        for fi33 := 0; fi33 < 19; fi33++ {
                            v34 := []int{32767, (fi33 | fi33), (fi33 * fi33)}
                            v25 = (v25 - (v25 % ((v25) | 1)))
                            switch (("hello" + (("" + "\n\t") + v3)) + v3) {
                            case "foo bar":
                                for fi35 := 0; fi35 < 20; fi35++ {
                                    v15 -= (^(v15))
                                    fmt.Println(v5)
                                    v25 = v25
                                    v36 := GenS4{F0: v8, F1: 1, F2: 86}
                                    _ = fi35
                                    _ = v36
                                }
                            case "hello":
                                v5 = v5
                            case "\n\t":
                                fmt.Println(v2, v8, v16)
                            }
                            _ = fi33
                            _ = v34
                        }
                    case "foo bar":
                        var v37 int8 = v1
                        _ = v37
                    }
                    _ = v25
                } else {
                    var v38 int16 = (v16 % ((v5) | 1))
                    _ = v38
                }
            }
            v39 := []int{-669, v15, -1}
            v40 := v4[1]
            _ = v39
            _ = v40
        } else {
            v2 = (v2 % ((v2) | 1))
        }
    }
    v7 = (-(v7))
    v8 = 42
    v41 := GenS5{F0: v5}
    for fi42 := 0; fi42 < 17; fi42++ {
        if false {
            v4[42] = int(v8)
        }
        var v43 int16 = -1
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
