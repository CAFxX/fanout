package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int64, p1 bool) int32 {
    fmt.Println(p0, p0)
    v4 := []int{-9223372036854775808, 9223372036854775807, -494, 0}
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v4
    return -532
}
func genFn1(p0 bool, p1 string, p2 uint32) uint16 {
    fmt.Println(p0, p1)
    for fi5 := 0; fi5 < 6; fi5++ {
        p1 = p1
        for fi6 := 0; fi6 < 20; fi6++ {
            p2 = p2
            if p0 {
                p0 = p0
                p1 = ((p1 + "\n\t") + p1)
            }
            _ = fi6
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint16 = 968
    v2 := []string{}
    var v3 int32 = -2147483648
    if len(v2) > 0 {
        v2[0] = ("\n\t" + "")
    }
    var v7 int16 = int16(v1)
    var v8 uint16 = v1
    var v9 uint8 = 127
    var v10 bool = (2147483647 != 42)
    if v10 {
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                v2 = append(v2, "a")
            } else {
                if (v7 > (v7 + 501)) {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        switch fi12 {
                        case -2147483648:
                            if (2 > 734) {
                                var v13 uint64 = 42
                                _ = v13
                            }
                            fmt.Println(v8, v7, v7)
                            fmt.Println(v3, v7, v9)
                            fmt.Println(v7, v10)
                            switch ("foo bar" + "foo bar") {
                            case "foo bar":
                                switch (("foo bar" + "foo bar") + "hello") {
                                case "foo bar":
                                    var v14 uint8 = v9
                                    _ = v14
                                case "a":
                                    var v15 float64 = (0.1 * 0.0)
                                    _ = v15
                                case "hello":
                                    fmt.Println(v3, v3, v10)
                                default:
                                    v1 = 65535
                                }
                            case "hello":
                                var v16 int = fi11
                                _ = v16
                            case "":
                                var v17 string = "\n\t"
                                _ = v17
                            default:
                                v8 -= (^(v8))
                            }
                        case 419:
                            _ = len(v2)
                        case 441:
                            var v18 int16 = v7
                            _ = v18
                        }
                        if (v10 && (v1 != (^(v8)))) {
                            fmt.Println(v7)
                            fmt.Println(v9)
                        }
                        _ = len(v2)
                        fmt.Println(fi12)
                        v10 = v10
                        v19 := GenS1{F0: 1}
                        if (4294967295 <= uint32(v1)) {
                            switch ("\n\t" + ("" + "\n\t")) {
                            case "\n\t":
                                v2 = append(v2, "\n\t")
                            case "a":
                                v20 := GenS1{F0: 867}
                                _ = v20
                            }
                            v1 = 42
                        } else {
                            _ = len(v2)
                            v21 := GenS1{F0: 42}
                            var v22 bool = v10
                            if len(v2) > 0 {
                                v2[0] = "foo bar"
                            }
                            var v23 uint32 = 65535
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = fi12
                        _ = v19
                    }
                }
            }
            v9 += (v9 / ((v9) | 1))
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(v1, fi11)
                if (0 <= v3) {
                    var v25 float32 = 0.1
                    v26 := GenS1{F0: 970}
                    _ = v25
                    _ = v26
                } else {
                    v3 = v3
                }
                var v27 string = ((((("" + ("foo bar" + "foo bar")) + "a") + "hello") + "foo bar") + ((("\n\t" + "foo bar") + "\n\t") + (("foo bar" + "foo bar") + "foo bar")))
                _ = fi24
                _ = v27
            }
            _ = fi11
        }
    }
    v3 -= (-(v3))
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 uint8 = (v9 ^ 61)
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
