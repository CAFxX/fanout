package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8 }

func main() {
    var v1 float64 = (-2.25 - 1e10)
    var v2 uint64 = 2
    var v3 float32 = ((1.5 * (((-2.25 + 0.0) + 0.1) - 0.1)) + (0.0 + 0.1))
    v4 := GenS1{F0: 69}
    var v5 int32 = 72
    var v6 uint8 = 57
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 42}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if (42 != uint(v2)) {
        v5 = v5
    } else {
        switch "foo bar" {
        case "foo bar":
            switch "foo bar" {
            case "a":
                v1 -= (float64(v3) * v1)
            case "hello":
                switch "" {
                case "hello":
                    if true {
                        v10 := GenS2{F0: 255}
                        v11 := GenS2{F0: 65535}
                        _ = v10
                        _ = v11
                    } else {
                        switch ("hello" + "a") {
                        case "hello":
                            fmt.Println(v6, v5, v6)
                        default:
                            var v12 uint32 = 302
                            _ = v12
                        }
                        v13 := GenS3{F0: true, F1: -1}
                        _ = v13
                    }
                case "foo bar":
                    v14 := GenS4{F0: ("foo bar" + "foo bar")}
                    _ = v14
                case "\n\t":
                    var v15 int16 = 1000
                    _ = v15
                default:
                    v16 := []int{}
                    _ = v16
                }
            case "\n\t":
                for fi17 := 0; fi17 < 5; fi17++ {
                    switch (fi17 + (fi17 - (fi17 & fi17))) {
                    case -32:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if ((!(100 > -1)) || true) {
                                v19 := GenS5{F0: -1}
                                v20 := []int{32767, (fi18 * 128), fi17}
                                v21 := GenS1{F0: 0}
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = fi18
                        }
                    }
                    switch -2147483648 {
                    case 682:
                        fmt.Println(v6, fi17, v1)
                    }
                    _ = fi17
                }
            }
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := []int{(-(fi22)), fi22, fi22, 716}
            if len(v23) > 0 {
                v23[0] = fi22
            }
            if len(v23) > 0 {
                v23[0] = fi22
            }
            _ = fi22
            _ = v23
        }
    }
    v24 := []int{9223372036854775807, -399}
    if true {
        v24 = append(v24, 2147483647)
        switch -145 {
        case -656:
            v25 := GenS3{F0: (!(!false)), F1: int(v2)}
            _ = v25
        case 32767:
            v6 = (v6 + 1)
        case 748:
            for fi26 := 0; fi26 < 10; fi26++ {
                if ("\n\t" <= "") {
                    var v27 uint32 = 42
                    _ = v27
                } else {
                    v28 := GenS4{F0: "a"}
                    v1 = v1
                    _ = v28
                }
                v29 := GenS3{F0: false, F1: fi26}
                switch ((("\n\t" + "\n\t") + ("foo bar" + "a")) + "\n\t") {
                case "":
                    var v30 int32 = v5
                    _ = v30
                }
                _ = fi26
                _ = v29
            }
        }
        var v31 uint8 = (^(v6))
        _ = v31
    } else {
        v6 = v6
        var v32 uint8 = (v6 << ((2) & 63))
        var v33 int8 = 1
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 255}
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 uint64 = 579
        _ = fi35
        _ = v36
    }
    var v37 int8 = -4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
