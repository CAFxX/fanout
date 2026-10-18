package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 int64, p1 int16, p2 int16) string {
    switch 2147483647 {
    case 32768:
        v5 := []int{}
        _ = v5
    case -675:
        var v6 uint32 = 852
        _ = v6
    case -2147483648:
        var v7 uint64 = 2
        _ = v7
    }
    var v8 int64 = p0
    v8 = (p0 << ((0) & 63))
    var v9 int16 = p2
    if (128 != 127) {
        var v10 int = 27
        if (!((v8 - (v8 - (p0 % ((v8) | 1)))) == v8)) {
            for fi11 := 0; fi11 < 12; fi11++ {
                v12 := GenS1{F0: 968}
                v10 += v10
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: 0}
                _ = fi13
                _ = v14
            }
            p2 += p2
        }
        var v15 uint32 = 65535
        _ = v10
        _ = v15
    } else {
        var v16 uint16 = 65535
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("foo bar" + ("foo bar" + ("\n\t" + ("foo bar" + ("a" + (((("foo bar" + "") + "a") + "\n\t") + ((("a" + "") + "a") + (("" + "foo bar") + "foo bar"))))))))
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = (1e10 - 0.0)
    var v3 uint8 = 197
    v4 := make(map[string]uint32)
    var v17 bool = true
    for fi18 := 0; fi18 < 4; fi18++ {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := []int{32768, fi19, (fi19 & fi18)}
            switch (fi19 / ((fi19) | 1)) {
            case 610:
                var v21 string = "hello"
                v17 = (((false || (42 < 65535)) || (v2 < v2)) && true)
                _ = v21
            case 793:
                if ((v1 / -2.25) <= 3.14159) {
                    var v22 uint64 = 18446744073709551615
                    _ = v22
                }
            }
            _ = fi19
            _ = v20
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := []int{(fi23 - (-(fi18))), (fi23 | fi18)}
            var v25 uint64 = 4294967295
            var v26 float32 = v2
            var v27 uint8 = v3
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 3; fi28++ {
            if v17 {
                var v29 bool = (65535 <= 42)
                _ = v29
            }
            _ = fi28
        }
        _ = fi18
    }
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "hello":
        if false {
            v30 := GenS1{F0: 1}
            v1 = v1
            _ = v30
        } else {
            var v31 uint16 = 904
            _ = v31
        }
    case "a":
        for fi32 := 0; fi32 < 16; fi32++ {
            _ = len(v4)
            if (-2147483648 < (fi32 * fi32)) {
                v3 += v3
                var v33 uint8 = (v3 << ((626) & 63))
                var v34 string = "\n\t"
                _ = v33
                _ = v34
            } else {
                if (!(1000 < 1)) {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        var v36 uint = 1
                        v37 := GenS2{F0: "\n\t"}
                        fmt.Println(v1, v3, v3)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                } else {
                    v1 += v1
                    delete(v4, "")
                    var v38 uint8 = 255
                    _ = v38
                }
            }
            var v39 uint64 = 42
            _ = fi32
            _ = v39
        }
    case "foo bar":
        for fi40 := 0; fi40 < 2; fi40++ {
            fmt.Println(v3, v2)
            var v41 int8 = 42
            _ = fi40
            _ = v41
        }
    default:
        v1 += v1
    }
    var v42 uint64 = 0
    v43 := GenS2{F0: "\n\t"}
    var v44 uint32 = 65535
    v45 := GenS2{F0: ""}
    var v46 uint8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
