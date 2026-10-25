package main

import (
    "fmt"
)

func genFn0(p0 int16, p1 uint8) int16 {
    p1 += p1
    var v6 int64 = 852
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    v1 := []string{"\n\t"}
    var v2 uint = 18446744073709551615
    v3 := make(map[string]uint32)
    var v4 string = "hello"
    var v5 int64 = 42
    var v7 int64 = v5
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 bool = true
        if (42 <= -991) {
            var v10 uint32 = 1
            v4 = "hello"
            v5 += v5
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    var v11 uint = 1
    for fi12 := 0; fi12 < 20; fi12++ {
        v11 = (v2 % ((v11) | 1))
        var v13 int16 = 1
        switch fi12 {
        case -748:
            var v14 int16 = v13
            _ = v14
        case -2147483648:
            switch (fi12 - (-(fi12))) {
            case 9223372036854775807:
                v7 -= (v7 % (((v7 / ((9223372036854775807) | 1))) | 1))
            case 4294967295:
                v15 := v3[(v4 + "\n\t")]
                _ = v15
            default:
                var v16 uint64 = 1
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                if (3.14159 <= (0.0 + 3.14159)) {
                    _ = len(v3)
                    v4 = ""
                    if true {
                        fmt.Println(fi12, v11)
                        v13 = (v13 - (v13 << ((v11) & 63)))
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi17
            }
        }
        _ = fi12
        _ = v13
    }
    var v18 float64 = 1e10
    fmt.Println(v2)
    var v19 bool = (255 < 32767)
    fmt.Println(v5, v18, v19)
    fmt.Println(v18, v5)
    var v20 float64 = float64(float32(float64((-2.25 - float32(v18)))))
    v7 = 9223372036854775807
    var v21 string = v4
    if ((-128 >= 0) || true) {
        v11 = (v2 + (v11 | (v2 / ((18446744073709551615) | 1))))
        var v22 string = "hello"
        var v23 bool = (!v19)
        _ = v22
        _ = v23
    }
    switch ("a" + "hello") {
    case "\n\t":
        var v24 int8 = -106
        _ = v24
    case "foo bar":
        var v25 uint64 = 2
        _ = v25
    case "":
        var v26 bool = (2147483647 == 1)
        _ = v26
    }
    v2 -= (v2 % (((v11 % ((v11) | 1))) | 1))
    var v27 int16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
