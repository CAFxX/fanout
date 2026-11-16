package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint8, p2 int32) int64 {
    p0 = p0
    if true {
        v7 := GenS1{F0: 144}
        p0 -= (p0 << ((0) & 63))
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := []int{}
            var v10 uint32 = 4294967295
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(p0, fi11)
        _ = fi11
    }
    v12 := []int{-1}
    var v13 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 9223372036854775807
}

func main() {
    v1 := []string{("a" + ("\n\t" + "a")), ("foo bar" + "\n\t"), ((("hello" + "hello") + ("hello" + "foo bar")) + "foo bar")}
    v2 := make(map[string]string)
    var v3 uint8 = 2
    v4 := make(map[string]bool)
    var v5 bool = (true || false)
    var v6 int = -317
    var v14 bool = v5
    v15 := GenS1{F0: 4294967295}
    if v5 {
        v16 := GenS1{F0: 79}
        _ = v16
    }
    delete(v4, "hello")
    v3 += (v3 | (v3 + (v3 >> ((484) & 63))))
    v6 = v6
    v5 = v14
    fmt.Println(v5, v5, v5)
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int8 = 127
        _ = fi17
        _ = v18
    }
    fmt.Println(v3, v5)
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 float32 = (1.5 + ((0.1 - (((1.5 / 1.5) + 3.14159) * 1.5)) - (3.14159 + -2.25)))
        if (float64(v20) <= 1.5) {
            v3 = 0
        } else {
            switch 127 {
            case -1:
                var v21 int8 = 42
                _ = v21
            case -354:
                v22 := GenS1{F0: 518}
                _ = v22
            }
        }
        var v23 string = "a"
        _ = fi19
        _ = v20
        _ = v23
    }
    var v24 string = "foo bar"
    for fi25 := 0; fi25 < 5; fi25++ {
        if (!((1e10 + 0.0) != 3.14159)) {
            for fi26 := 0; fi26 < 19; fi26++ {
                if v5 {
                    var v27 int = (-(v6))
                    var v28 uint32 = 2147483648
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v27
                    _ = v28
                } else {
                    fmt.Println(fi26, v24)
                    fmt.Println(v6, v5)
                }
                v24 = v24
                var v29 uint16 = 984
                _ = fi26
                _ = v29
            }
            v30 := GenS1{F0: 247}
            _ = v30
        }
        var v31 uint32 = 0
        v5 = v5
        _ = fi25
        _ = v31
    }
    if v14 {
        v24 = v24
        for fi32 := 0; fi32 < 7; fi32++ {
            v3 += (^(v3))
            _ = fi32
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            v5 = (2147483647 <= -802)
            _ = fi33
        }
    }
    v6 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
