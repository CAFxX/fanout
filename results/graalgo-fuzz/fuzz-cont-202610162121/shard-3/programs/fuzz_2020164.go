package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
func genFn0() uint64 {
    var v6 float64 = (1e10 * 0.1)
    v7 := GenS1{F0: (1 >= 32767)}
    for fi8 := 0; fi8 < 12; fi8++ {
        if (("hello" + "foo bar") <= ("a" + "foo bar")) {
            v6 += ((3.14159 + (float64(float32(float64(0.0))) * (v6 + float64(3.14159)))) + 3.14159)
        } else {
            v6 += 0.0
        }
        _ = fi8
    }
    var v9 int64 = 11
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            if true {
                var v12 uint8 = 255
                if false {
                    switch "foo bar" {
                    case "hello":
                        var v13 int16 = 1000
                        _ = v13
                    case "\n\t":
                        if false {
                            v14 := []int{(fi11 / ((fi11) | 1)), fi11}
                            v15 := GenS1{F0: false}
                            _ = v14
                            _ = v15
                        }
                    }
                    v6 += 0.0
                }
                v9 = v9
                _ = v12
            } else {
                fmt.Println(v6)
                var v16 int = (-(fi10))
                var v17 bool = true
                _ = v16
                _ = v17
            }
            var v18 uint64 = 42
            _ = fi11
            _ = v18
        }
        _ = fi10
    }
    _ = v6
    _ = v7
    _ = v9
    return 95
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    var v3 int64 = -9223372036854775808
    var v4 int16 = 42
    var v5 int64 = 9223372036854775807
    v19 := []int{32767, 42, 127, 9223372036854775807}
    fmt.Println(v4, v5, v3)
    _ = len(v2)
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            for fi21 := 0; fi21 < 11; fi21++ {
                if true {
                    var v22 float32 = (0.0 - 1e10)
                    _ = v22
                }
                var v23 float64 = 1.5
                _ = fi21
                _ = v23
            }
        } else {
            if (true && v1) {
                fmt.Println(v1, v1, v3)
                fmt.Println(v4)
            } else {
                var v24 bool = (!(!true))
                switch fi20 {
                case 32767:
                    fmt.Println(v3, v3)
                case 885:
                    for fi25 := 0; fi25 < 16; fi25++ {
                        _ = len(v19)
                        switch fi25 {
                        case 1:
                            fmt.Println(v1)
                        case 42:
                            v4 += (v4 & (v4 * (v4 % ((v4) | 1))))
                        default:
                            v3 -= 0
                        }
                        var v26 float64 = -2.25
                        _ = fi25
                        _ = v26
                    }
                case 2147483647:
                    for fi27 := 0; fi27 < 3; fi27++ {
                        v28 := GenS1{F0: v1}
                        _ = fi27
                        _ = v28
                    }
                }
                _ = v24
            }
        }
        _ = fi20
    }
    if (false && (719 == 1)) {
        v5 += -1
    }
    v3 = 1
    for fi29 := 0; fi29 < 6; fi29++ {
        v5 = (v3 >> ((874) & 63))
        _ = fi29
    }
    switch -197 {
    case -644:
        fmt.Println(v4)
    case -296:
        if v1 {
            v4 = (v4 - v4)
            v30 := GenS1{F0: v1}
            _ = v30
        }
    default:
        v1 = ((!(488 > 1)) || false)
    }
    var v31 float64 = 0.1
    var v32 int64 = v3
    v19 = append(v19, 127)
    fmt.Println(v32, v4)
    var v33 uint64 = 122
    v34 := v2[("foo bar" + "\n\t")]
    v35 := v2[(("\n\t" + "\n\t") + ((("" + "a") + "foo bar") + "foo bar"))]
    v36 := GenS2{F0: 42}
    var v37 string = "a"
    if v1 {
        for fi38 := 0; fi38 < 15; fi38++ {
            for fi39 := 0; fi39 < 19; fi39++ {
                v37 = "a"
                v1 = false
                _ = fi39
            }
            switch ("" + v37) {
            case "hello":
                var v40 bool = (false || (v35 == (fi38 ^ v34)))
                _ = v40
            }
            var v41 int = (v35 & v35)
            _ = fi38
            _ = v41
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
