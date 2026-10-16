package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 string; F2 uint16 }
func genFn0(p0 int32) uint64 {
    if (634 < 954) {
        p0 = p0
        v5 := []int{-988, -429, 42}
        _ = v5
    }
    var v6 bool = (float32(0.0) < 1e10)
    v7 := []int{32768}
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int32) string {
    v8 := GenS1{F0: 2, F1: ("" + ("hello" + "")), F2: (42 >= 42)}
    switch 42 {
    case -2147483648:
        p0 -= p0
    case 32768:
        var v9 uint = 255
        _ = v9
    }
    fmt.Println(p0)
    p0 = 0
    _ = p0
    _ = v8
    return ""
}

func main() {
    var v1 bool = (42 < 2)
    var v2 int = 127
    var v3 uint32 = 65535
    v4 := make(map[string]string)
    fmt.Println(v2)
    fmt.Println(v3, v2, v2)
    if true {
        if (1 <= (v2 >> ((65535) & 63))) {
            v2 += (v2 >> ((255) & 63))
        } else {
            v10 := GenS2{F0: ""}
            switch v2 {
            case 9223372036854775807:
                var v11 int8 = -1
                _ = v11
            case 127:
                var v12 float64 = ((0.1 / 0.1) / float64(-2.25))
                _ = v12
            default:
                if v1 {
                    v1 = (!("foo bar" < ("hello" + ("a" + "a"))))
                    v2 -= v2
                } else {
                    v13 := GenS3{F0: 42}
                    var v14 bool = (-2.25 == ((-2.25 - 3.14159) * ((1e10 + float64((((-2.25 + float32(1.5)) * (float32(1.5) / 3.14159)) - 1e10))) + (3.14159 + (0.0 / 0.1)))))
                    _ = v13
                    _ = v14
                }
            }
            _ = v10
        }
        v15 := GenS4{F0: v3}
        _ = v15
    } else {
        var v16 uint16 = 32767
        v17 := GenS3{F0: -497}
        var v18 int64 = 434
        _ = v16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            switch (fi20 / ((127) | 1)) {
            case 690:
                switch (fi19 ^ (^(v2))) {
                case -380:
                    fmt.Println(v2, v1, fi19)
                }
                var v21 float64 = 3.14159
                v1 = v1
                if (42 <= 390) {
                    v3 = 2
                    v22 := []int{v2}
                    v23 := GenS4{F0: v3}
                    fmt.Println(v2, fi19, v1)
                    v3 += 2147483648
                    v3 = 1
                    _ = v22
                    _ = v23
                }
                v1 = (!false)
                fmt.Println(v21, v21)
                _ = v21
            case -130:
                fmt.Println(fi19, v2)
            }
            var v24 int64 = 9223372036854775807
            var v25 uint32 = 42
            switch fi20 {
            case 9223372036854775807:
                v26 := GenS5{F0: 0}
                _ = v26
            case 786:
                var v27 bool = v1
                _ = v27
            default:
                v28 := []int{(v2 | fi19), (-(fi20)), fi19}
                _ = v28
            }
            v3 = (v3 - (v25 | 2147483647))
            _ = fi20
            _ = v24
            _ = v25
        }
        v1 = true
        for fi29 := 0; fi29 < 1; fi29++ {
            fmt.Println(v2, v2, v1)
            v2 = fi29
            _ = fi29
        }
        _ = fi19
    }
    v4["hello"] = "\n\t"
    v3 += v3
    var v30 uint8 = 1
    if ((!true) || ((1 > 32767) || ((v2 - (v2 << ((375) & 63))) >= v2))) {
        delete(v4, "hello")
        for fi31 := 0; fi31 < 3; fi31++ {
            var v32 uint8 = v30
            var v33 int32 = 42
            _ = fi31
            _ = v32
            _ = v33
        }
    } else {
        v34 := []int{-391, v2, (^(v2))}
        _ = v34
    }
    v1 = v1
    var v35 int = 2147483647
    fmt.Println(v2)
    var v36 int64 = 1
    if (!false) {
        v37 := GenS1{F0: 65535, F1: ("\n\t" + "hello"), F2: v1}
        _ = v37
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 bool = (446 >= 962)
        switch -2147483648 {
        case 128:
            if (0.0 <= (3.14159 - (-2.25 + -2.25))) {
                v40 := GenS4{F0: v3}
                if (v1 && (float32(1e10) != (0.1 + 0.0))) {
                    v41 := GenS6{F0: ((v39 && v39) && v1), F1: "a", F2: 798}
                    _ = v41
                } else {
                    v42 := GenS3{F0: (-(v36))}
                    _ = v42
                }
                _ = v40
            }
        case 4294967295:
            fmt.Println(v30)
        case 530:
            v35 = 2147483647
        }
        var v43 uint = uint(v30)
        _ = fi38
        _ = v39
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
}
