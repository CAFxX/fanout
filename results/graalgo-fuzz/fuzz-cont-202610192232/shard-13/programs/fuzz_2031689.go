package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint64 {
    var v7 int32 = 2147483647
    var v8 uint64 = 4294967295
    if false {
        v9 := GenS1{F0: 933}
        v8 = (v8 >> ((18446744073709551615) & 63))
        v10 := []int{127, 0, 32767}
        _ = v10
        _ = v9
    } else {
        switch "" {
        case "foo bar":
            switch -349 {
            case 32768:
                v8 += (^(v8))
            case 42:
                v11 := []int{0, 127}
                _ = v11
            case 9223372036854775807:
                var v12 string = "hello"
                _ = v12
            default:
                var v13 uint = 4294967295
                _ = v13
            }
        case "\n\t":
            var v14 float32 = ((-2.25 * 3.14159) * 1e10)
            _ = v14
        }
    }
    var v15 float32 = (((float32(float64((1.5 + float32(0.0)))) - 0.0) + (float32((float64(0.1) + 1e10)) / ((1.5 * -2.25) - (float32(1.5) - -2.25)))) - 0.0)
    v7 = (v7 << ((123) & 63))
    _ = v15
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 uint64, p1 int) int16 {
    if (42 >= 1) {
        p1 += (p1 >> ((552) & 63))
        p0 -= 4294967295
        for fi16 := 0; fi16 < 19; fi16++ {
            p0 = p0
            var v17 int64 = 9223372036854775807
            v18 := GenS1{F0: fi16}
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            v20 := []int{}
            var v21 float32 = 1e10
            var v22 int64 = 9223372036854775807
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 bool = (!true)
            _ = fi23
            _ = v24
        }
        fmt.Println(p1, p0)
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        p0 = genFn0()
        fmt.Println(p1, p1, p0)
        if false {
            if true {
                v26 := []int{127, fi25, (fi25 + -2147483648)}
                if (128 < 42) {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        _ = len(v26)
                        fmt.Println(fi25, fi25)
                        _ = fi27
                    }
                }
                _ = v26
            }
            v28 := []int{fi25, 0, -779}
            var v29 uint32 = uint32(p0)
            _ = v28
            _ = v29
        }
        _ = fi25
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 bool = true
    var v2 int64 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[string]uint32)
    v6 := []int{524, 2147483647, -518}
    var v30 uint16 = 42
    if false {
        for fi31 := 0; fi31 < 16; fi31++ {
            for fi32 := 0; fi32 < 16; fi32++ {
                v6 = append(v6, 32768)
                var v33 uint16 = 0
                _ = fi32
                _ = v33
            }
            if len(v6) > 0 {
                v6[0] = fi31
            }
            _ = fi31
        }
        var v34 string = ("hello" + "\n\t")
        _ = v34
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 int = (fi35 << ((255) & 63))
        _ = fi35
        _ = v36
    }
    v37 := GenS1{F0: -9223372036854775808}
    fmt.Println(v1, v30)
    v1 = (!(100 == 42))
    var v38 int = 0
    v2 = 0
    if (-791 >= 1000) {
        var v39 int8 = -1
        v30 -= 2
        _ = v39
    } else {
        switch "hello" {
        case "hello":
            v2 = -1
        case "":
            v38 = (v38 / ((v38) | 1))
        }
        delete(v4, 32767)
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (0 >= 18446744073709551615) {
                var v40 int16 = -32768
                _ = v40
            } else {
                fmt.Println(v30, v2)
                var v41 bool = v1
                _ = v41
            }
        } else {
            var v42 uint = 0
            v43 := GenS1{F0: (v38 ^ (v38 ^ 127))}
            _ = v42
            _ = v43
        }
    }
    delete(v3, (v38 ^ v38))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
