package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 string; F1 string; F2 int8 }

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 4294967295
    v3 := make(map[uint32]string)
    var v4 uint = 459
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := []int{32767, (fi5 % ((127) | 1)), fi5}
        v7 := GenS1{F0: 1}
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        v2 = v2
        v3[1] = v1
    } else {
        if (42 < 106) {
            if ((v4 ^ (v4 & 2)) <= v4) {
                var v8 int16 = -349
                for fi9 := 0; fi9 < 11; fi9++ {
                    switch "a" {
                    case "a":
                        for fi10 := 0; fi10 < 2; fi10++ {
                            var v11 int64 = 9223372036854775807
                            v12 := []int{(-(fi9)), 1, fi9}
                            v13 := GenS1{F0: uint64(v2)}
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    }
                    _ = fi9
                }
                _ = v8
            } else {
                var v14 bool = false
                v14 = false
                var v15 int16 = 983
                _ = v14
                _ = v15
            }
            if true {
                var v16 bool = (!true)
                v16 = (!v16)
                v17 := []int{42}
                _ = v16
                _ = v17
            } else {
                v2 -= (v2 >> (((v4 << ((v4) & 63))) & 63))
                switch (v1 + ("\n\t" + v1)) {
                case "a":
                    if true {
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 uint64 = 859
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        if false {
                            v20 := GenS1{F0: 314}
                            var v21 float32 = (1.5 * (0.1 + float32((0.1 / (-2.25 / ((1.5 + (3.14159 * -2.25)) + 0.1))))))
                            var v22 float32 = 3.14159
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            v4 = (v4 / ((v4) | 1))
                        }
                        var v23 int = 371
                        v24 := GenS2{F0: -1, F1: false, F2: 763}
                        _ = v23
                        _ = v24
                    }
                case "foo bar":
                    v4 -= 0
                case "":
                    v1 = v1
                }
            }
            v25 := GenS3{F0: 0}
            _ = v25
        }
        var v26 uint8 = 127
        for fi27 := 0; fi27 < 7; fi27++ {
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := GenS4{F0: (^(v26)), F1: 1}
                v4 -= (v4 | v4)
                _ = fi28
                _ = v29
            }
            var v30 int64 = -9223372036854775808
            for fi31 := 0; fi31 < 7; fi31++ {
                if ((((^(v30)) < v30) || (false && true)) || ((^(v2)) == 65535)) {
                    v32 := GenS5{F0: (("foo bar" + v1) + v1), F1: v1, F2: -27}
                    v33 := []int{(-(fi27)), fi27}
                    _ = v32
                    _ = v33
                }
                _ = fi31
            }
            _ = fi27
            _ = v30
        }
        _ = v26
    }
    if (1.5 >= 1.5) {
        fmt.Println(v4, v1)
        v34 := []int{2147483647}
        _ = v34
    }
    var v35 float64 = (1e10 - (-2.25 * (((3.14159 * (0.0 / (-2.25 - 3.14159))) - 1e10) + (0.0 - (3.14159 + -2.25)))))
    v36 := GenS1{F0: 42}
    for fi37 := 0; fi37 < 4; fi37++ {
        v2 = 42
        v35 -= ((1e10 + v35) / float64(1e10))
        _ = fi37
    }
    if (!false) {
        switch (v1 + v1) {
        case "hello":
            var v38 int64 = -9223372036854775808
            _ = v38
        default:
            var v39 int16 = 982
            _ = v39
        }
    } else {
        if true {
            for fi40 := 0; fi40 < 10; fi40++ {
                v41 := GenS4{F0: 2, F1: 448}
                for fi42 := 0; fi42 < 2; fi42++ {
                    v3[(v2 | v2)] = "\n\t"
                    v35 = 1e10
                    _ = fi42
                }
                _ = fi40
                _ = v41
            }
        }
        delete(v3, v2)
    }
    fmt.Println(v4, v1, v35)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
