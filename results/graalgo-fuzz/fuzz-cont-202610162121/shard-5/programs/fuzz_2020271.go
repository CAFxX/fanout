package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
func genFn0(p0 string) bool {
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = -2.25
        switch "foo bar" {
        case "":
            for fi8 := 0; fi8 < 2; fi8++ {
                switch fi8 {
                case -623:
                    v9 := []int{(fi6 / ((fi6) | 1))}
                    _ = v9
                default:
                    v10 := []int{fi6}
                    if false {
                        var v11 int32 = 0
                        _ = v11
                    } else {
                        if (!(2 == 42)) {
                            var v12 int16 = -32768
                            _ = v12
                        }
                        var v13 uint16 = 1
                        var v14 uint64 = 674
                        _ = v13
                        _ = v14
                    }
                    _ = v10
                }
                var v15 bool = (!false)
                var v16 bool = (1.5 != (0.1 * (1e10 / 0.1)))
                _ = fi8
                _ = v15
                _ = v16
            }
        default:
            fmt.Println(v7, p0, fi6)
        }
        _ = fi6
        _ = v7
    }
    v17 := []int{32768, 32768, 128}
    if len(v17) > 0 {
        v17[0] = -9223372036854775808
    }
    v18 := GenS1{F0: p0, F1: 1, F2: 112}
    _ = p0
    _ = v17
    _ = v18
    return false
}

func main() {
    var v1 int8 = 1
    var v2 float64 = 0.0
    var v3 float64 = (0.1 - v2)
    var v4 float32 = 0.1
    var v5 int64 = 42
    switch ((("\n\t" + ("" + "foo bar")) + "foo bar") + "a") {
    case "\n\t":
        var v19 bool = ((!false) || true)
        _ = v19
    }
    v4 = 3.14159
    var v20 float32 = v4
    for fi21 := 0; fi21 < 13; fi21++ {
        switch fi21 {
        case 2147483647:
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 uint32 = 337
                var v24 float64 = 0.0
                v24 += v3
                _ = fi22
                _ = v23
                _ = v24
            }
        case 105:
            v3 += float64(v4)
        case 1:
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 int32 = -2147483648
                if (163 <= 42) {
                    if (42 > -32768) {
                        v20 -= float32(float64(0.0))
                    }
                    v27 := []int{2147483647, (fi25 << ((65535) & 63)), fi25}
                    _ = v27
                } else {
                    var v28 uint = 18446744073709551615
                    if false {
                        var v29 float64 = v2
                        _ = v29
                    }
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        }
        v5 = 345
        _ = fi21
    }
    var v30 int32 = -1
    if true {
        var v31 int8 = v1
        _ = v31
    }
    var v32 int8 = (v1 * (v1 + 42))
    v33 := GenS2{F0: ("a" + "a"), F1: 17}
    var v34 bool = (false || (1 >= 4294967295))
    var v35 bool = ((v20 * ((v4 - 1.5) + -2.25)) < v4)
    var v36 float32 = 3.14159
    if false {
        var v37 uint64 = 42
        _ = v37
    }
    if (true && v35) {
        for fi38 := 0; fi38 < 19; fi38++ {
            v39 := GenS3{F0: 4294967295, F1: 1, F2: 255}
            v40 := []int{fi38, fi38, (-(fi38)), fi38}
            var v41 uint = 1
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    }
    var v42 uint8 = 255
    v43 := []int{4294967295, -2147483648}
    var v44 int64 = -566
    switch ("\n\t" + ("hello" + ("a" + ("a" + ("hello" + "a"))))) {
    case "a":
        v44 += 0
    case "foo bar":
        v30 = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
