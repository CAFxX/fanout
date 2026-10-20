package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    var v1 string = ("" + "hello")
    var v2 uint8 = 255
    var v3 uint8 = v2
    var v4 uint16 = 255
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(v4)
        if true {
            v3 = v3
            if (!((float64(1e10) * -2.25) < 3.14159)) {
                var v6 int32 = 1
                fmt.Println(v1, fi5)
                v7 := GenS1{F0: (v4 ^ (v4 + v4)), F1: 4294967295, F2: v2}
                _ = v6
                _ = v7
            }
        } else {
            var v8 uint = 2
            var v9 uint32 = 1
            v2 = v3
            _ = v8
            _ = v9
        }
        var v10 int8 = 42
        _ = fi5
        _ = v10
    }
    v3 = 2
    v1 = ("hello" + "\n\t")
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int32 = 420
        if (-662 < -9223372036854775808) {
            var v13 bool = (true || (2 > 65535))
            v14 := []int{fi11, fi11, (fi11 - (fi11 % ((fi11) | 1))), 116}
            switch v1 {
            case "\n\t":
                if v13 {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                        v16 := GenS1{F0: 1, F1: 2147483647, F2: v3}
                        var v17 int = (-(fi11))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    v1 = "foo bar"
                } else {
                    switch (fi11 & 32768) {
                    case -627:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            if len(v14) > 0 {
                                v14[0] = (^(fi11))
                            }
                            fmt.Println(fi18)
                            var v19 int16 = 1000
                            _ = fi18
                            _ = v19
                        }
                    case 60:
                        if len(v14) > 0 {
                            v14[0] = 112
                        }
                    }
                    var v20 float64 = -2.25
                    _ = v20
                }
            case "a":
                if len(v14) > 0 {
                    v14[0] = (fi11 << ((42) & 63))
                }
            }
            _ = v13
            _ = v14
        } else {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 int32 = v12
                v22 += (v22 - v22)
                var v23 int8 = 0
                var v24 int = 2147483647
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 uint8 = 128
            if false {
                switch "foo bar" {
                case "foo bar":
                    for fi26 := 0; fi26 < 9; fi26++ {
                        if true {
                            v27 := GenS1{F0: (v4 % (((v4 & v4)) | 1)), F1: 65535, F2: (v3 / (((^(v3))) | 1))}
                            if true {
                                v3 -= 255
                            } else {
                                if true {
                                    v28 := GenS1{F0: v4, F1: 2147483648, F2: (v25 & (v2 / ((v2) | 1)))}
                                    fmt.Println(v1, v1)
                                    if (42 >= 42) {
                                        v29 := []int{-139, (^(fi11)), fi26}
                                        _ = v29
                                    }
                                    _ = v28
                                } else {
                                    var v30 bool = ((fi11 | (fi11 ^ 2147483647)) != fi11)
                                    v31 := GenS1{F0: (^(v4)), F1: 2147483647, F2: v3}
                                    v25 = (^(v2))
                                    _ = v30
                                    _ = v31
                                }
                                v32 := []int{427, fi26}
                                _ = v32
                            }
                            _ = v27
                        }
                        _ = fi26
                    }
                case "hello":
                    v2 = v25
                case "":
                    var v33 uint32 = 0
                    _ = v33
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    v25 -= 1
                    v1 = v1
                    _ = fi34
                }
                v35 := GenS1{F0: 1, F1: 697, F2: v3}
                _ = v35
            }
            _ = v25
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    switch 0 {
    case -194:
        if (0 != 579) {
            v4 += v4
        }
    case 0:
        var v36 bool = (!false)
        _ = v36
    }
    v37 := []int{32768, 157, 32768, -9223372036854775808}
    var v38 bool = (2147483647 == 744)
    if v38 {
        for fi39 := 0; fi39 < 16; fi39++ {
            var v40 int64 = -9223372036854775808
            _ = fi39
            _ = v40
        }
    }
    v38 = v38
    if (4294967295 >= uint64(v4)) {
        var v41 string = "a"
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
