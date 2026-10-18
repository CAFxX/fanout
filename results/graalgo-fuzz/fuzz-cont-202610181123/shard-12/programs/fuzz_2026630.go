package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 uint64; F2 int64 }
type GenS5 struct { F0 int }

func main() {
    var v1 int16 = 248
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 int16 = (v1 & (v1 / ((v1) | 1)))
    fmt.Println(v4, v1)
    v1 -= (-(v4))
    v5 := []int{32767, -2147483648, -958, 303}
    if true {
        v6 := GenS1{F0: ("a" + ("" + "\n\t"))}
        v7 := GenS1{F0: "a"}
        v8 := GenS2{F0: 0, F1: 0, F2: true}
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS3{F0: ""}
    var v10 int = 128
    v4 -= -32768
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(fi11)
        switch 0 {
        case 2147483647:
            if (false && false) {
                v12 := GenS1{F0: "a"}
                _ = v12
            } else {
                v13 := GenS1{F0: "foo bar"}
                fmt.Println(v10, fi11)
                _ = v13
            }
        case -719:
            switch ("hello" + "\n\t") {
            case "a":
                switch (^(fi11)) {
                case 4294967295:
                    var v14 int32 = -2147483648
                    _ = v14
                }
            case "\n\t":
                switch (v10 + v10) {
                case 0:
                    switch v10 {
                    case 127:
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v2[65535] = fi11
                            _ = fi15
                        }
                    case 2147483647:
                        switch "a" {
                        case "hello":
                            var v16 uint8 = 255
                            _ = v16
                        case "\n\t":
                            switch "" {
                            case "\n\t":
                                v4 = (v1 >> ((636) & 63))
                            }
                        case "a":
                            v4 += v1
                        }
                    case -333:
                        var v17 string = "a"
                        _ = v17
                    }
                case 2147483647:
                    for fi18 := 0; fi18 < 15; fi18++ {
                        if (true && (true || true)) {
                            v19 := GenS4{F0: 42, F1: 724, F2: -9223372036854775808}
                            if false {
                                v20 := GenS4{F0: -128, F1: 1, F2: 75}
                                v21 := GenS4{F0: -10, F1: 1, F2: -37}
                                var v22 int = fi18
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            _ = v19
                        } else {
                            var v23 float64 = -2.25
                            _ = v23
                        }
                        v5 = append(v5, int(v1))
                        _ = fi18
                    }
                }
            case "":
                for fi24 := 0; fi24 < 3; fi24++ {
                    fmt.Println(v1, v1, v4)
                    v25 := GenS5{F0: (fi24 ^ fi24)}
                    var v26 int8 = -128
                    if (!(("foo bar" + ("foo bar" + ("foo bar" + "a"))) > "")) {
                        _ = len(v5)
                        var v27 uint = 1
                        _ = v27
                    }
                    v28 := GenS2{F0: (fi24 + (fi11 | 32767)), F1: 646, F2: (32767 <= 42)}
                    _ = len(v2)
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v28
                }
            }
        case 234:
            var v29 bool = (!true)
            _ = v29
        }
        v30 := GenS5{F0: (fi11 | v10)}
        _ = fi11
        _ = v30
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
