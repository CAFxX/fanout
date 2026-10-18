package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0() int16 {
    var v6 uint16 = 32767
    v6 = v6
    switch "a" {
    case "hello":
        switch "hello" {
        case "\n\t":
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := GenS1{F0: 831}
                _ = fi7
                _ = v8
            }
        case "hello":
            v6 -= 32767
        case "":
            v6 += (v6 & v6)
        default:
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v9 int16 = -1
                _ = v9
            case 42:
                var v10 int8 = -1
                _ = v10
            case 128:
                for fi11 := 0; fi11 < 11; fi11++ {
                    for fi12 := 0; fi12 < 2; fi12++ {
                        if ((1.5 - 3.14159) != float64((0.0 / 0.1))) {
                            v13 := []int{fi12, fi11}
                            _ = v13
                        } else {
                            var v14 uint16 = v6
                            _ = v14
                        }
                        v6 = v6
                        v6 = v6
                        fmt.Println(fi12, v6, fi11)
                        _ = fi12
                    }
                    _ = fi11
                }
            default:
                v6 += v6
            }
        }
    }
    _ = v6
    return 1
}

func main() {
    var v1 uint8 = 2
    var v2 bool = (!(!true))
    v3 := make(map[int]string)
    v4 := []string{"hello"}
    v5 := []int{2147483647, 42}
    switch "a" {
    case "foo bar":
        for fi15 := 0; fi15 < 18; fi15++ {
            v2 = (!(!v2))
            v1 += 237
            _ = fi15
        }
    case "a":
        v16 := GenS1{F0: 0}
        _ = v16
    default:
        var v17 int16 = 42
        _ = v17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        for fi18 := 0; fi18 < 20; fi18++ {
            v3[fi18] = (((("hello" + (("" + "") + "foo bar")) + "") + "foo bar") + ("\n\t" + ""))
            v1 -= v1
            var v19 int32 = 122
            _ = fi18
            _ = v19
        }
    case "a":
        v20 := GenS2{F0: 942, F1: "\n\t"}
        _ = v20
    case "":
        var v21 int16 = 695
        _ = v21
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint64 = 18446744073709551615
        v24 := GenS2{F0: (v23 + 4294967295), F1: (("a" + "") + "hello")}
        fmt.Println(v23)
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 uint16 = 65535
    fmt.Println(v2, v2, v25)
    if true {
        v1 = v1
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS1{F0: 42}
            v1 += v1
            _ = fi26
            _ = v27
        }
    } else {
        if v2 {
            switch 1 {
            case -133:
                v28 := GenS3{F0: 1, F1: (v25 - 255)}
                _ = v28
            case -2147483648:
                var v29 int = -1
                _ = v29
            case -9223372036854775808:
                var v30 float32 = (-2.25 + 1e10)
                _ = v30
            default:
                if (!(v2 || ((v25 | 65535) == (v25 / (((^(v25))) | 1))))) {
                    v4 = append(v4, "foo bar")
                    var v31 int64 = int64(v25)
                    _ = v31
                }
            }
            fmt.Println(v2, v1, v1)
        }
        v25 = (v25 ^ (v25 ^ (v25 ^ 1)))
        for fi32 := 0; fi32 < 10; fi32++ {
            if len(v5) > 0 {
                v5[0] = fi32
            }
            v1 = (^(v1))
            if (725 != 65535) {
                v2 = ((v2 || (!(false || v2))) || false)
                var v33 uint = 929
                _ = v33
            }
            _ = fi32
        }
    }
    var v34 int8 = 42
    var v35 int8 = 100
    var v36 float64 = (1.5 + 0.0)
    var v37 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
