package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
func genFn0() string {
    if false {
        v6 := GenS1{F0: 32767}
        var v7 int32 = 1
        v7 = v7
        _ = v6
        _ = v7
    }
    v8 := []int{2147483647, 127, 60, 4294967295}
    var v9 int32 = -732
    for fi10 := 0; fi10 < 16; fi10++ {
        v9 += (v9 * v9)
        v9 -= v9
        _ = fi10
    }
    _ = v8
    _ = v9
    return ("a" + (("\n\t" + "") + ""))
}

func main() {
    var v1 uint64 = 42
    var v2 bool = (!(139 == 233))
    v3 := make(map[string]string)
    var v4 int = 42
    v5 := []string{}
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v1)
            var v13 uint16 = 316
            _ = fi12
            _ = v13
        }
        if (v2 || (v2 && (-877 != -817))) {
            v2 = ("hello" != "hello")
            switch (^(fi11)) {
            case -633:
                switch fi11 {
                case 4294967295:
                    fmt.Println(fi11)
                case 127:
                    switch "" {
                    case "a":
                        v1 = 259
                    case "foo bar":
                        if (!v2) {
                            if (("foo bar" == "a") && v2) {
                                if len(v5) > 0 {
                                    v5[0] = (("\n\t" + ("a" + "")) + "foo bar")
                                }
                                fmt.Println(fi11, v2, v4)
                                var v14 uint8 = 123
                                _ = v14
                            }
                            fmt.Println(v1)
                        } else {
                            var v15 bool = (v1 >= 42)
                            var v16 uint32 = 42
                            for fi17 := 0; fi17 < 4; fi17++ {
                                _ = len(v3)
                                fmt.Println(fi11)
                                fmt.Println(v2)
                                _ = fi17
                            }
                            _ = v15
                            _ = v16
                        }
                    case "\n\t":
                        for fi18 := 0; fi18 < 1; fi18++ {
                            if ((^(v1)) <= v1) {
                                delete(v3, "a")
                                v1 += v1
                                v19 := GenS1{F0: fi11}
                                _ = v19
                            }
                            _ = fi18
                        }
                    }
                }
            case -1:
                switch fi11 {
                case 9223372036854775807:
                    fmt.Println(v1)
                case -756:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        switch (fi20 - (-(fi11))) {
                        case 0:
                            var v21 bool = false
                            _ = v21
                        case -385:
                            fmt.Println(fi20)
                        default:
                            if v2 {
                                fmt.Println(fi11)
                                if (v2 || v2) {
                                    switch v4 {
                                    case 884:
                                        v4 = -1
                                    default:
                                        v4 = -1
                                    }
                                    var v22 bool = (((28 >= 0) || v2) && v2)
                                    _ = v22
                                }
                            }
                        }
                        _ = fi20
                    }
                case 127:
                    var v23 int16 = 1000
                    _ = v23
                default:
                    var v24 uint32 = 2147483648
                    _ = v24
                }
            case 127:
                for fi25 := 0; fi25 < 11; fi25++ {
                    var v26 uint64 = 814
                    _ = fi25
                    _ = v26
                }
            }
        } else {
            for fi27 := 0; fi27 < 17; fi27++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v28 string = "foo bar"
                _ = fi27
                _ = v28
            }
        }
        if (!v2) {
            v29 := v3["a"]
            var v30 uint64 = v1
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 13; fi31++ {
                v32 := GenS2{F0: 2147483647, F1: 4294967295, F2: int16(v4)}
                var v33 uint64 = v1
                v1 = 0
                _ = fi31
                _ = v32
                _ = v33
            }
            v1 = v1
            var v34 uint16 = 2
            _ = v34
        }
        _ = fi11
    }
    for fi35 := 0; fi35 < 14; fi35++ {
        var v36 int = 4294967295
        _ = fi35
        _ = v36
    }
    fmt.Println(v2)
    v2 = v2
    var v37 uint32 = 2147483647
    var v38 bool = (!false)
    v1 = v1
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 19; fi40++ {
            v4 += (fi40 / ((-1) | 1))
            if v2 {
                v4 += (v4 | fi39)
                fmt.Println(v38, v4, v4)
                if v38 {
                    var v41 uint16 = 162
                    _ = v41
                } else {
                    var v42 uint32 = (v37 << ((uint(v37)) & 63))
                    _ = v42
                }
            }
            var v43 uint16 = 1
            _ = fi40
            _ = v43
        }
        var v44 int8 = 1
        _ = fi39
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
