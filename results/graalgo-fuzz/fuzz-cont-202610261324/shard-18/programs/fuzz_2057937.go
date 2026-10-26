package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint16) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        if (1e10 >= 1e10) {
            var v6 bool = false
            _ = v6
        } else {
            if true {
                p0 -= 42
            } else {
                v7 := []int{-1, fi5, (fi5 / ((fi5) | 1)), fi5}
                var v8 uint64 = 42
                v9 := GenS1{F0: "\n\t", F1: fi5}
                _ = v7
                _ = v8
                _ = v9
            }
            v10 := []int{fi5, (fi5 & fi5), 32767, (fi5 / ((fi5) | 1))}
            _ = v10
        }
        _ = fi5
    }
    var v11 string = ""
    var v12 int = 32768
    if ((3.14159 * (((((0.1 - 0.0) * (float64(float32(0.1)) * ((float64(0.1) + float64(3.14159)) - 1.5))) * 1e10) * float64((float32(0.0) / -2.25))) / -2.25)) <= (0.0 - (1e10 * 0.1))) {
        v13 := GenS2{F0: 2}
        var v14 string = v11
        fmt.Println(v12, v14)
        _ = v13
        _ = v14
    } else {
        var v15 uint64 = 1
        var v16 uint32 = 337
        var v17 bool = (18446744073709551615 == 424)
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := []int{(v12 / (((v12 * v12)) | 1)), (v12 * v12)}
    _ = p0
    _ = v11
    _ = v12
    _ = v18
    return true
}
func genFn1() uint8 {
    for fi19 := 0; fi19 < 13; fi19++ {
        fmt.Println(fi19)
        v20 := GenS2{F0: 18446744073709551615}
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 542, F1: -1, F2: 65535}
    v22 := []int{}
    var v23 int8 = 84
    var v24 int64 = -1
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 int16 = 552
    var v2 int16 = -620
    var v3 int32 = 0
    v4 := make(map[string]int)
    v4[("\n\t" + "\n\t")] = -416
    var v25 string = "hello"
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := []int{}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        v1 -= v2
        _ = fi28
    }
    v3 -= (v3 >> ((1) & 63))
    if false {
        switch -214 {
        case -2147483648:
            fmt.Println(v1)
        case -550:
            var v29 int8 = -1
            _ = v29
        case 2147483647:
            v3 = -1
        }
        var v30 int64 = -1
        v31 := GenS4{F0: 372, F1: 177, F2: 866}
        _ = v30
        _ = v31
    } else {
        v2 += (v1 - v1)
    }
    if (float32((0.1 + ((0.0 / (0.0 - (1.5 * ((-2.25 / float64(3.14159)) + (1e10 - 1e10))))) * -2.25))) == 0.0) {
        var v32 uint64 = 1
        _ = v32
    } else {
        v2 = (v2 << ((18446744073709551615) & 63))
        v33 := []int{352, -952, 9223372036854775807}
        var v34 string = ((("hello" + (((v25 + v25) + "") + v25)) + (v25 + (v25 + ("foo bar" + (v25 + v25))))) + ("a" + ((v25 + v25) + "\n\t")))
        _ = v33
        _ = v34
    }
    var v35 uint64 = 42
    if false {
        v25 = (v25 + "foo bar")
    } else {
        for fi36 := 0; fi36 < 7; fi36++ {
            switch (fi36 << ((42) & 63)) {
            case 9223372036854775807:
                delete(v4, v25)
            }
            _ = fi36
        }
    }
    var v37 bool = true
    v38 := GenS4{F0: 1, F1: 42, F2: 65535}
    v3 = (v3 - v3)
    switch 1 {
    case 907:
        fmt.Println(v25, v3)
    case 4294967295:
        fmt.Println(v25)
    case 1:
        for fi39 := 0; fi39 < 10; fi39++ {
            v1 = -785
            var v40 float32 = -2.25
            v41 := GenS5{F0: v37}
            _ = fi39
            _ = v40
            _ = v41
        }
    default:
        v42 := GenS3{F0: 672, F1: 0, F2: 65535}
        _ = v42
    }
    v25 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v4
}
