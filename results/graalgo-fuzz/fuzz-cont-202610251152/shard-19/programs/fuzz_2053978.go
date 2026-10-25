package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int64 = 513
    v2 := make(map[string]bool)
    var v3 bool = false
    v4 := []string{"foo bar", "a", "foo bar"}
    var v5 uint = 65535
    var v6 int16 = 22
    v3 = v3
    v5 += (v5 >> ((v5) & 63))
    if (1.5 < ((1e10 / float64(-2.25)) - 0.0)) {
        v5 = v5
    } else {
        v1 = (v1 & 9223372036854775807)
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 bool = ((((v5 < (v5 >> ((v5) & 63))) && v3) && (!(0.1 < 1.5))) && (!v3))
            for fi9 := 0; fi9 < 3; fi9++ {
                fmt.Println(v3, fi7, v8)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    v2[("\n\t" + ("hello" + "foo bar"))] = false
    v3 = (9223372036854775807 == v1)
    var v10 uint64 = 145
    if v3 {
        v5 = (^(v5))
    } else {
        for fi11 := 0; fi11 < 20; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 bool = (!v3)
                v10 -= 4294967295
                _ = fi12
                _ = v13
            }
            switch ("\n\t" + "hello") {
            case "":
                for fi14 := 0; fi14 < 7; fi14++ {
                    if (42 > 18446744073709551615) {
                        if v3 {
                            fmt.Println(v1)
                        } else {
                            fmt.Println(v1, fi11)
                        }
                    }
                    if (v3 && ((1.5 + float64((1.5 + 1e10))) == 0.0)) {
                        if (-1 <= -1) {
                            _ = len(v2)
                            var v15 uint = 134
                            switch fi11 {
                            case 9223372036854775807:
                                delete(v2, (("hello" + "foo bar") + ("" + "a")))
                            }
                            _ = v15
                        }
                        var v16 float32 = 1.5
                        _ = v16
                    } else {
                        v5 = v5
                        v1 -= v1
                        v17 := GenS1{F0: (v1 + (v1 - v1))}
                        if true {
                            v18 := v2["foo bar"]
                            switch "a" {
                            case "\n\t":
                                v10 = v10
                            case "foo bar":
                                v4 = append(v4, (("\n\t" + "") + "foo bar"))
                            default:
                                var v19 string = "\n\t"
                                _ = v19
                            }
                            v20 := GenS1{F0: -1}
                            _ = v18
                            _ = v20
                        }
                        fmt.Println(fi11, fi11)
                        _ = v17
                    }
                    _ = fi14
                }
            case "\n\t":
                var v21 int16 = v6
                _ = v21
            default:
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 uint = v5
                    fmt.Println(v10)
                    var v24 int64 = v1
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = fi11
        }
    }
    fmt.Println(v3, v1)
    v5 += 822
    var v25 int64 = (v1 % (((^(v1))) | 1))
    v26 := v2[""]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
