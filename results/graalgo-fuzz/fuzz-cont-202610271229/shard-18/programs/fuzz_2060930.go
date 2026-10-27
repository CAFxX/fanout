package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 string) int {
    v7 := []int{}
    p0 = (p0 + "hello")
    _ = p0
    _ = v7
    return 314
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 205
    v3 := make(map[int]int)
    var v4 int16 = 32767
    v5 := []uint32{42}
    var v6 uint16 = 255
    switch "" {
    case "a":
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 int = 120
                _ = fi9
                _ = v10
            }
            if ((0.0 == -2.25) || (!true)) {
                var v11 int64 = 916
                v4 = v4
                _ = v11
            } else {
                if (!false) {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v4 = v4
                        var v13 uint = 18446744073709551615
                        v1 = append(v1, 2147483648)
                        _ = fi12
                        _ = v13
                    }
                }
                var v14 float64 = 1.5
                switch (^(fi8)) {
                case 4294967295:
                    var v15 uint16 = v6
                    _ = v15
                }
                _ = v14
            }
            _ = fi8
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = v6
    v16 := GenS1{F0: -1, F1: (true || false)}
    v4 = v4
    for fi17 := 0; fi17 < 10; fi17++ {
        if ((v2 / ((v2) | 1)) >= 102) {
            var v18 uint8 = (v2 >> ((4294967295) & 63))
            if false {
                if ((v6 % (((^(v6))) | 1)) > v6) {
                    for fi19 := 0; fi19 < 13; fi19++ {
                        var v20 uint8 = (v2 | (v2 - (v18 ^ 255)))
                        _ = fi19
                        _ = v20
                    }
                    var v21 float32 = 1.5
                    _ = v21
                }
            }
            v2 = v2
            _ = v18
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi17
    }
    var v22 uint64 = 20
    v23 := v3[4294967295]
    if ((42 <= 2147483647) || ((("foo bar" + ((("" + ("" + "hello")) + (("\n\t" + "foo bar") + "")) + "")) + "a") != "\n\t")) {
        for fi24 := 0; fi24 < 16; fi24++ {
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := GenS1{F0: 1, F1: false}
                switch "foo bar" {
                case "\n\t":
                    v6 -= 775
                default:
                    v6 -= 65535
                }
                _ = fi25
                _ = v26
            }
            var v27 int64 = -1
            _ = fi24
            _ = v27
        }
        if true {
            switch (v23 << ((18446744073709551615) & 63)) {
            case -642:
                v4 = v4
            case 558:
                fmt.Println(v2)
            }
        } else {
            if (v6 == v6) {
                v28 := GenS1{F0: 127, F1: (true || false)}
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 float32 = -2.25
                    if false {
                        if (4294967295 >= 1) {
                            v4 += 32767
                        }
                        for fi31 := 0; fi31 < 5; fi31++ {
                            v32 := GenS2{F0: 18446744073709551615, F1: v22, F2: (v6 + v6)}
                            switch (^(fi29)) {
                            case 128:
                                var v33 int = v23
                                _ = v33
                            case 127:
                                v22 = (v22 + 18446744073709551615)
                            }
                            _ = fi31
                            _ = v32
                        }
                    } else {
                        var v34 uint32 = 0
                        var v35 uint8 = v2
                        _ = v34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 18; fi36++ {
                        v37 := GenS1{F0: 42, F1: true}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v38 string = "\n\t"
                        var v39 string = ((("\n\t" + v38) + v38) + "foo bar")
                        _ = fi36
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = fi29
                    _ = v30
                }
                _ = v28
            }
        }
        if len(v1) > 0 {
            v1[0] = 4294967295
        }
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        v41 := GenS3{F0: (^(fi40)), F1: false}
        var v42 float32 = 3.14159
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
